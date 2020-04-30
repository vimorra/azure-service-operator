# Image URL to use all building/pushing image targets
IMG ?= controller:latest

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"

BUILD_ID ?= $(shell git rev-parse --short HEAD)

# best to keep the prefix as short as possible to not exceed naming limits for things like keyvault (24 chars)
TEST_RESOURCE_PREFIX ?= aso-$(BUILD_ID)

# Go compiler builds tags: some parts of the test suite use these to selectively compile tests.
BUILD_TAGS ?= all

all: manager

# Generate test certs for development
generate-test-certs:
	echo "[req]" > config.txt
	echo "distinguished_name = req_distinguished_name" >> config.txt
	echo "[req_distinguished_name]" >> config.txt
	echo "[SAN]" >> config.txt
	echo "subjectAltName=DNS:azureoperator-webhook-service.azureoperator-system.svc.cluster.local" >> config.txt
	openssl req -x509 -days 730 -out tls.crt -keyout tls.key -newkey rsa:4096 -subj "/CN=azureoperator-webhook-service.azureoperator-system" -config config.txt -nodes
	rm -rf /tmp/k8s-webhook-server
	mkdir -p /tmp/k8s-webhook-server/serving-certs
	mv tls.* /tmp/k8s-webhook-server/serving-certs/

# Run Controller tests against the configured cluster
test-integration-controllers: generate fmt vet manifests
	TEST_RESOURCE_PREFIX=$(TEST_RESOURCE_PREFIX) TEST_USE_EXISTING_CLUSTER=true REQUEUE_AFTER=20 \
	go test -v -tags "$(BUILD_TAGS)" -coverprofile=reports/integration-controllers-coverage-output.txt -coverpkg=./... -covermode count -parallel 4 -timeout 45m \
	./controllers/... \
	2>&1 | tee reports/integration-controllers-output.txt
	go-junit-report < reports/integration-controllers-output.txt > reports/integration-controllers-report.xml

# Run Resource Manager tests against the configured cluster
test-integration-managers: generate fmt vet manifests
	TEST_USE_EXISTING_CLUSTER=true TEST_CONTROLLER_WITH_MOCKS=false REQUEUE_AFTER=20 \
	go test -v -coverprofile=reports/integration-managers-coverage-ouput.txt -coverpkg=./... -covermode count -parallel 4 -timeout 45m \
	./api/... \
	./pkg/resourcemanager/eventhubs/...  \
	./pkg/resourcemanager/resourcegroups/...  \
	./pkg/resourcemanager/storages/... \
	./pkg/resourcemanager/psql/server/... \
	./pkg/resourcemanager/psql/database/... \
	./pkg/resourcemanager/psql/firewallrule/... \
	./pkg/resourcemanager/appinsights/... \
	./pkg/resourcemanager/vnet/...
	2>&1 | tee reports/integration-managers-output.txt
	go-junit-report < reports/integration-managers-output.txt > reports/integration-managers-report.xml

# Run all available tests. Note that Controllers are not unit-testable.
test-unit: generate fmt vet manifests 
	TEST_USE_EXISTING_CLUSTER=false REQUEUE_AFTER=20 \
	go test -v -tags "$(BUILD_TAGS)" -coverprofile=coverage-unit.txt -covermode count -parallel 4 -timeout 10m \
	./api/... \
	./pkg/secrets/...
	./pkg/resourcemanager/keyvaults/unittest/ \
	2>&1 | tee testlogs.txt
	go-junit-report < testlogs.txt > report-unit.xml
	go tool cover -html=coverage/coverage.txt -o cover-unit.html

# Merge all the available test coverage results and publish a single report
test-process-coverage:
	find reports -name "*-coverage-output.txt" -type f -print | xargs gocovmerge > reports/merged-coverage-output.txt
	gocov convert reports/merged-coverage-output.txt > reports/merged-coverage-output.json
	gocov-xml < reports/merged-coverage-output.json > reports/merged-coverage.xml
	go tool cover -html=reports/merged-coverage-output.txt -o reports/merged-coverage.html

# Cleanup resource groups azure created by tests using pattern matching 't-rg-'
test-cleanup-azure-resources: 	
	# Delete the resource groups that match the pattern
	for rgname in `az group list --query "[*].[name]" -o table | grep '^${TEST_RESOURCE_PREFIX}' `; do \
	    echo "$$rgname will be deleted"; \
	    az group delete --name $$rgname --no-wait --yes; \
    done

# Build manager binary
manager: generate fmt vet
	go build -o manager main.go

# Run against the configured Kubernetes cluster in ~/.kube/config
run: generate fmt vet
	go run ./main.go

# Install CRDs into a cluster
install: generate
	kubectl apply -f config/crd/bases

# Deploy controller in the configured Kubernetes cluster in ~/.kube/config
deploy: manifests
	kubectl apply -f config/crd/bases
	kustomize build config/default | kubectl apply -f -

timestamp := $(shell /bin/date "+%Y%m%d-%H%M%S")

update:
	IMG="docker.io/controllertest:$(timestamp)" make ARGS="${ARGS}" docker-build
	kind load docker-image docker.io/controllertest:$(timestamp) --loglevel "trace"
	make install
	make deploy
	sed -i'' -e 's@image: .*@image: '"IMAGE_URL"'@' ./config/default/manager_image_patch.yaml

delete:
	kubectl delete -f config/crd/bases
	kustomize build config/default | kubectl delete -f -

# Validate copyright headers
validate-copyright-headers:
	@./scripts/validate-copyright-headers.sh

# Generate manifests for helm and package them up
helm-chart-manifests: manifests
	mkdir charts/azure-service-operator/templates/generated
	kustomize build ./config/default -o ./charts/azure-service-operator/templates/generated
	rm charts/azure-service-operator/templates/generated/~g_v1_namespace_azureoperator-system.yaml
	sed -i '' -e 's@controller:latest@{{ .Values.image.repository }}@' ./charts/azure-service-operator/templates/generated/apps_v1_deployment_azureoperator-controller-manager.yaml
	find ./charts/azure-service-operator/templates/generated/ -type f -exec sed -i '' -e 's@namespace: azureoperator-system@namespace: {{ .Values.namespace }}@' {} \;
	helm package ./charts/azure-service-operator -d ./charts
	helm repo index ./charts

delete-helm-gen-manifests:
	rm -rf charts/azure-service-operator/templates/generated/

# Generate manifests e.g. CRD, RBAC etc.
manifests: controller-gen
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet: 
	go vet ./...

# Generate code
generate: manifests
	$(CONTROLLER_GEN) object:headerFile=./hack/boilerplate.go.txt paths=./api/...

# find or download controller-gen
# download controller-gen if necessary
controller-gen:
ifeq (, $(shell which controller-gen))
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.0
CONTROLLER_GEN=$(shell go env GOPATH)/bin/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

.PHONY: install-bindata
install-bindata:
	go get -u github.com/jteeuwen/go-bindata/...

.PHONE:
generate-template:
	go-bindata -pkg template -prefix pkg/template/assets/ -o pkg/template/templates.go pkg/template/assets/

install-kubebuilder:
ifeq (,$(shell which kubebuilder))
	@echo "installing kubebuilder"
	# download kubebuilder and extract it to tmp
	curl -sL https://go.kubebuilder.io/dl/2.0.0/$(shell go env GOOS)/$(shell go env GOARCH) | tar -xz -C /tmp/
	# move to a long-term location and put it on your path
	# (you'll need to set the KUBEBUILDER_ASSETS env var if you put it somewhere else)
	mv /tmp/kubebuilder_2.0.0_$(shell go env GOOS)_$(shell go env GOARCH) /usr/local/kubebuilder
	export PATH=$$PATH:/usr/local/kubebuilder/bin
else
	@echo "kubebuilder has been installed"
endif

install-kustomize:
ifeq (,$(shell which kustomize))
	@echo "installing kustomize"
	mkdir -p /usr/local/kubebuilder/bin
	# download kustomize
	curl -o /usr/local/kubebuilder/bin/kustomize -sL "https://go.kubebuilder.io/kustomize/$(shell go env GOOS)/$(shell go env GOARCH)"
	# set permission
	chmod a+x /usr/local/kubebuilder/bin/kustomize
	$(shell which kustomize)
else
	@echo "kustomize has been installed"
endif

install-aad-pod-identity:
	kubectl apply -f https://raw.githubusercontent.com/Azure/aad-pod-identity/master/deploy/infra/deployment-rbac.yaml

install-test-dependencies:
	go get github.com/jstemmer/go-junit-report \
	&& go get github.com/axw/gocov/gocov \
	&& go get github.com/AlekSi/gocov-xml \
	&& go get github.com/wadey/gocovmerge
