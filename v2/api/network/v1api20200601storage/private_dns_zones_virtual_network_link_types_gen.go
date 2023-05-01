// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601storage

import (
	v1api20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=privatednszonesvirtualnetworklinks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={privatednszonesvirtualnetworklinks/status,privatednszonesvirtualnetworklinks/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20200601.PrivateDnsZonesVirtualNetworkLink
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}
type PrivateDnsZonesVirtualNetworkLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsZones_VirtualNetworkLink_Spec   `json:"spec,omitempty"`
	Status            PrivateDnsZones_VirtualNetworkLink_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateDnsZonesVirtualNetworkLink{}

// GetConditions returns the conditions of the resource
func (link *PrivateDnsZonesVirtualNetworkLink) GetConditions() conditions.Conditions {
	return link.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (link *PrivateDnsZonesVirtualNetworkLink) SetConditions(conditions conditions.Conditions) {
	link.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &PrivateDnsZonesVirtualNetworkLink{}

// AzureName returns the Azure name of the resource
func (link *PrivateDnsZonesVirtualNetworkLink) AzureName() string {
	return link.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (link PrivateDnsZonesVirtualNetworkLink) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (link *PrivateDnsZonesVirtualNetworkLink) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (link *PrivateDnsZonesVirtualNetworkLink) GetSpec() genruntime.ConvertibleSpec {
	return &link.Spec
}

// GetStatus returns the status of this resource
func (link *PrivateDnsZonesVirtualNetworkLink) GetStatus() genruntime.ConvertibleStatus {
	return &link.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones/virtualNetworkLinks"
func (link *PrivateDnsZonesVirtualNetworkLink) GetType() string {
	return "Microsoft.Network/privateDnsZones/virtualNetworkLinks"
}

// NewEmptyStatus returns a new empty (blank) status
func (link *PrivateDnsZonesVirtualNetworkLink) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateDnsZones_VirtualNetworkLink_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (link *PrivateDnsZonesVirtualNetworkLink) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(link.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  link.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (link *PrivateDnsZonesVirtualNetworkLink) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateDnsZones_VirtualNetworkLink_STATUS); ok {
		link.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateDnsZones_VirtualNetworkLink_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	link.Status = st
	return nil
}

// Hub marks that this PrivateDnsZonesVirtualNetworkLink is the hub type for conversion
func (link *PrivateDnsZonesVirtualNetworkLink) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (link *PrivateDnsZonesVirtualNetworkLink) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: link.Spec.OriginalVersion,
		Kind:    "PrivateDnsZonesVirtualNetworkLink",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20200601.PrivateDnsZonesVirtualNetworkLink
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/virtualNetworkLinks/{virtualNetworkLinkName}
type PrivateDnsZonesVirtualNetworkLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsZonesVirtualNetworkLink `json:"items"`
}

// Storage version of v1api20200601.PrivateDnsZones_VirtualNetworkLink_Spec
type PrivateDnsZones_VirtualNetworkLink_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Etag            *string `json:"etag,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/PrivateDnsZone resource
	Owner               *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"PrivateDnsZone"`
	PropertyBag         genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RegistrationEnabled *bool                              `json:"registrationEnabled,omitempty"`
	Tags                map[string]string                  `json:"tags,omitempty"`
	VirtualNetwork      *SubResource                       `json:"virtualNetwork,omitempty"`
}

var _ genruntime.ConvertibleSpec = &PrivateDnsZones_VirtualNetworkLink_Spec{}

// ConvertSpecFrom populates our PrivateDnsZones_VirtualNetworkLink_Spec from the provided source
func (link *PrivateDnsZones_VirtualNetworkLink_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == link {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(link)
}

// ConvertSpecTo populates the provided destination from our PrivateDnsZones_VirtualNetworkLink_Spec
func (link *PrivateDnsZones_VirtualNetworkLink_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == link {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(link)
}

// Storage version of v1api20200601.PrivateDnsZones_VirtualNetworkLink_STATUS
type PrivateDnsZones_VirtualNetworkLink_STATUS struct {
	Conditions              []conditions.Condition `json:"conditions,omitempty"`
	Etag                    *string                `json:"etag,omitempty"`
	Id                      *string                `json:"id,omitempty"`
	Location                *string                `json:"location,omitempty"`
	Name                    *string                `json:"name,omitempty"`
	PropertyBag             genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState       *string                `json:"provisioningState,omitempty"`
	RegistrationEnabled     *bool                  `json:"registrationEnabled,omitempty"`
	Tags                    map[string]string      `json:"tags,omitempty"`
	Type                    *string                `json:"type,omitempty"`
	VirtualNetwork          *SubResource_STATUS    `json:"virtualNetwork,omitempty"`
	VirtualNetworkLinkState *string                `json:"virtualNetworkLinkState,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PrivateDnsZones_VirtualNetworkLink_STATUS{}

// ConvertStatusFrom populates our PrivateDnsZones_VirtualNetworkLink_STATUS from the provided source
func (link *PrivateDnsZones_VirtualNetworkLink_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == link {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(link)
}

// ConvertStatusTo populates the provided destination from our PrivateDnsZones_VirtualNetworkLink_STATUS
func (link *PrivateDnsZones_VirtualNetworkLink_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == link {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(link)
}

// Storage version of v1api20200601.SubResource
// Reference to another subresource.
type SubResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// AssignProperties_From_SubResource populates our SubResource from the provided source SubResource
func (resource *SubResource) AssignProperties_From_SubResource(source *v1api20201101s.SubResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Reference
	if source.Reference != nil {
		reference := source.Reference.Copy()
		resource.Reference = &reference
	} else {
		resource.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SubResource populates the provided destination SubResource from our SubResource
func (resource *SubResource) AssignProperties_To_SubResource(destination *v1api20201101s.SubResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Reference
	if resource.Reference != nil {
		reference := resource.Reference.Copy()
		destination.Reference = &reference
	} else {
		destination.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20200601.SubResource_STATUS
// Reference to another subresource.
type SubResource_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_SubResource_STATUS populates our SubResource_STATUS from the provided source SubResource_STATUS
func (resource *SubResource_STATUS) AssignProperties_From_SubResource_STATUS(source *v1api20201101s.SubResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource_STATUS); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SubResource_STATUS populates the provided destination SubResource_STATUS from our SubResource_STATUS
func (resource *SubResource_STATUS) AssignProperties_To_SubResource_STATUS(destination *v1api20201101s.SubResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSubResource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSubResource_STATUS); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSubResource interface {
	AssignPropertiesFrom(src *v1api20201101s.SubResource) error
	AssignPropertiesTo(dst *v1api20201101s.SubResource) error
}

type augmentConversionForSubResource_STATUS interface {
	AssignPropertiesFrom(src *v1api20201101s.SubResource_STATUS) error
	AssignPropertiesTo(dst *v1api20201101s.SubResource_STATUS) error
}

func init() {
	SchemeBuilder.Register(&PrivateDnsZonesVirtualNetworkLink{}, &PrivateDnsZonesVirtualNetworkLinkList{})
}