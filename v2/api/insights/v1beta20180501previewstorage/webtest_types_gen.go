// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180501previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=insights.azure.com,resources=webtests,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=insights.azure.com,resources={webtests/status,webtests/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20180501preview.Webtest
// Generated from: https://schema.management.azure.com/schemas/2018-05-01-preview/Microsoft.Insights.Application.json#/resourceDefinitions/webtests
type Webtest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Webtests_Spec  `json:"spec,omitempty"`
	Status            WebTest_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Webtest{}

// GetConditions returns the conditions of the resource
func (webtest *Webtest) GetConditions() conditions.Conditions {
	return webtest.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (webtest *Webtest) SetConditions(conditions conditions.Conditions) {
	webtest.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &Webtest{}

// AzureName returns the Azure name of the resource
func (webtest *Webtest) AzureName() string {
	return webtest.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01-preview"
func (webtest Webtest) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (webtest *Webtest) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (webtest *Webtest) GetSpec() genruntime.ConvertibleSpec {
	return &webtest.Spec
}

// GetStatus returns the status of this resource
func (webtest *Webtest) GetStatus() genruntime.ConvertibleStatus {
	return &webtest.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/webtests"
func (webtest *Webtest) GetType() string {
	return "Microsoft.Insights/webtests"
}

// NewEmptyStatus returns a new empty (blank) status
func (webtest *Webtest) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &WebTest_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (webtest *Webtest) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(webtest.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  webtest.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (webtest *Webtest) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*WebTest_Status); ok {
		webtest.Status = *st
		return nil
	}

	// Convert status to required version
	var st WebTest_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	webtest.Status = st
	return nil
}

// Hub marks that this Webtest is the hub type for conversion
func (webtest *Webtest) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (webtest *Webtest) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: webtest.Spec.OriginalVersion,
		Kind:    "Webtest",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20180501preview.Webtest
// Generated from: https://schema.management.azure.com/schemas/2018-05-01-preview/Microsoft.Insights.Application.json#/resourceDefinitions/webtests
type WebtestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webtest `json:"items"`
}

// +kubebuilder:validation:Enum={"2018-05-01-preview"}
type APIVersion string

const APIVersionValue = APIVersion("2018-05-01-preview")

// Storage version of v1beta20180501preview.WebTest_Status
type WebTest_Status struct {
	Conditions         []conditions.Condition                    `json:"conditions,omitempty"`
	Configuration      *WebTestProperties_Status_Configuration   `json:"Configuration,omitempty"`
	Description        *string                                   `json:"Description,omitempty"`
	Enabled            *bool                                     `json:"Enabled,omitempty"`
	Frequency          *int                                      `json:"Frequency,omitempty"`
	Id                 *string                                   `json:"id,omitempty"`
	Kind               *string                                   `json:"Kind,omitempty"`
	Location           *string                                   `json:"location,omitempty"`
	Locations          []WebTestGeolocation_Status               `json:"Locations,omitempty"`
	Name               *string                                   `json:"name,omitempty"`
	PropertiesName     *string                                   `json:"properties_name,omitempty"`
	PropertyBag        genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	ProvisioningState  *string                                   `json:"provisioningState,omitempty"`
	Request            *WebTestProperties_Status_Request         `json:"Request,omitempty"`
	RetryEnabled       *bool                                     `json:"RetryEnabled,omitempty"`
	SyntheticMonitorId *string                                   `json:"SyntheticMonitorId,omitempty"`
	Tags               *v1.JSON                                  `json:"tags,omitempty"`
	Timeout            *int                                      `json:"Timeout,omitempty"`
	Type               *string                                   `json:"type,omitempty"`
	ValidationRules    *WebTestProperties_Status_ValidationRules `json:"ValidationRules,omitempty"`
}

var _ genruntime.ConvertibleStatus = &WebTest_Status{}

// ConvertStatusFrom populates our WebTest_Status from the provided source
func (test *WebTest_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == test {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(test)
}

// ConvertStatusTo populates the provided destination from our WebTest_Status
func (test *WebTest_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == test {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(test)
}

// Storage version of v1beta20180501preview.Webtests_Spec
type Webtests_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                          `json:"azureName,omitempty"`
	Configuration   *WebTestPropertiesConfiguration `json:"Configuration,omitempty"`
	Description     *string                         `json:"Description,omitempty"`
	Enabled         *bool                           `json:"Enabled,omitempty"`
	Frequency       *int                            `json:"Frequency,omitempty"`
	Kind            *string                         `json:"Kind,omitempty"`
	Location        *string                         `json:"location,omitempty"`
	Locations       []WebTestGeolocation            `json:"Locations,omitempty"`
	Name            *string                         `json:"Name,omitempty"`
	OriginalVersion string                          `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner              *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag        genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Request            *WebTestPropertiesRequest          `json:"Request,omitempty"`
	RetryEnabled       *bool                              `json:"RetryEnabled,omitempty"`
	SyntheticMonitorId *string                            `json:"SyntheticMonitorId,omitempty"`
	Tags               map[string]string                  `json:"tags,omitempty"`
	Timeout            *int                               `json:"Timeout,omitempty"`
	ValidationRules    *WebTestPropertiesValidationRules  `json:"ValidationRules,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Webtests_Spec{}

// ConvertSpecFrom populates our Webtests_Spec from the provided source
func (webtests *Webtests_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == webtests {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(webtests)
}

// ConvertSpecTo populates the provided destination from our Webtests_Spec
func (webtests *Webtests_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == webtests {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(webtests)
}

// Storage version of v1beta20180501preview.WebTestGeolocation
// Generated from: https://schema.management.azure.com/schemas/2018-05-01-preview/Microsoft.Insights.Application.json#/definitions/WebTestGeolocation
type WebTestGeolocation struct {
	Id          *string                `json:"Id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20180501preview.WebTestGeolocation_Status
type WebTestGeolocation_Status struct {
	Id          *string                `json:"Id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20180501preview.WebTestPropertiesConfiguration
// Generated from: https://schema.management.azure.com/schemas/2018-05-01-preview/Microsoft.Insights.Application.json#/definitions/WebTestPropertiesConfiguration
type WebTestPropertiesConfiguration struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WebTest     *string                `json:"WebTest,omitempty"`
}

// Storage version of v1beta20180501preview.WebTestPropertiesRequest
// Generated from: https://schema.management.azure.com/schemas/2018-05-01-preview/Microsoft.Insights.Application.json#/definitions/WebTestPropertiesRequest
type WebTestPropertiesRequest struct {
	FollowRedirects        *bool                  `json:"FollowRedirects,omitempty"`
	Headers                []HeaderField          `json:"Headers,omitempty"`
	HttpVerb               *string                `json:"HttpVerb,omitempty"`
	ParseDependentRequests *bool                  `json:"ParseDependentRequests,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestBody            *string                `json:"RequestBody,omitempty"`
	RequestUrl             *string                `json:"RequestUrl,omitempty"`
}

// Storage version of v1beta20180501preview.WebTestPropertiesValidationRules
// Generated from: https://schema.management.azure.com/schemas/2018-05-01-preview/Microsoft.Insights.Application.json#/definitions/WebTestPropertiesValidationRules
type WebTestPropertiesValidationRules struct {
	ContentValidation             *WebTestPropertiesValidationRulesContentValidation `json:"ContentValidation,omitempty"`
	ExpectedHttpStatusCode        *int                                               `json:"ExpectedHttpStatusCode,omitempty"`
	IgnoreHttpsStatusCode         *bool                                              `json:"IgnoreHttpsStatusCode,omitempty"`
	PropertyBag                   genruntime.PropertyBag                             `json:"$propertyBag,omitempty"`
	SSLCertRemainingLifetimeCheck *int                                               `json:"SSLCertRemainingLifetimeCheck,omitempty"`
	SSLCheck                      *bool                                              `json:"SSLCheck,omitempty"`
}

// Storage version of v1beta20180501preview.WebTestProperties_Status_Configuration
type WebTestProperties_Status_Configuration struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	WebTest     *string                `json:"WebTest,omitempty"`
}

// Storage version of v1beta20180501preview.WebTestProperties_Status_Request
type WebTestProperties_Status_Request struct {
	FollowRedirects        *bool                  `json:"FollowRedirects,omitempty"`
	Headers                []HeaderField_Status   `json:"Headers,omitempty"`
	HttpVerb               *string                `json:"HttpVerb,omitempty"`
	ParseDependentRequests *bool                  `json:"ParseDependentRequests,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequestBody            *string                `json:"RequestBody,omitempty"`
	RequestUrl             *string                `json:"RequestUrl,omitempty"`
}

// Storage version of v1beta20180501preview.WebTestProperties_Status_ValidationRules
type WebTestProperties_Status_ValidationRules struct {
	ContentValidation             *WebTestProperties_Status_ValidationRules_ContentValidation `json:"ContentValidation,omitempty"`
	ExpectedHttpStatusCode        *int                                                        `json:"ExpectedHttpStatusCode,omitempty"`
	IgnoreHttpsStatusCode         *bool                                                       `json:"IgnoreHttpsStatusCode,omitempty"`
	PropertyBag                   genruntime.PropertyBag                                      `json:"$propertyBag,omitempty"`
	SSLCertRemainingLifetimeCheck *int                                                        `json:"SSLCertRemainingLifetimeCheck,omitempty"`
	SSLCheck                      *bool                                                       `json:"SSLCheck,omitempty"`
}

// Storage version of v1beta20180501preview.HeaderField
// Generated from: https://schema.management.azure.com/schemas/2018-05-01-preview/Microsoft.Insights.Application.json#/definitions/HeaderField
type HeaderField struct {
	Key         *string                `json:"key,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20180501preview.HeaderField_Status
type HeaderField_Status struct {
	Key         *string                `json:"key,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20180501preview.WebTestPropertiesValidationRulesContentValidation
// Generated from: https://schema.management.azure.com/schemas/2018-05-01-preview/Microsoft.Insights.Application.json#/definitions/WebTestPropertiesValidationRulesContentValidation
type WebTestPropertiesValidationRulesContentValidation struct {
	ContentMatch    *string                `json:"ContentMatch,omitempty"`
	IgnoreCase      *bool                  `json:"IgnoreCase,omitempty"`
	PassIfTextFound *bool                  `json:"PassIfTextFound,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20180501preview.WebTestProperties_Status_ValidationRules_ContentValidation
type WebTestProperties_Status_ValidationRules_ContentValidation struct {
	ContentMatch    *string                `json:"ContentMatch,omitempty"`
	IgnoreCase      *bool                  `json:"IgnoreCase,omitempty"`
	PassIfTextFound *bool                  `json:"PassIfTextFound,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Webtest{}, &WebtestList{})
}
