// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	computev1alpha1api20201201 "github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20201201"
	"github.com/Azure/azure-service-operator/v2/api/compute/v1alpha1api20201201storage"
	computev1beta20201201 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201"
	"github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type VirtualMachineExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *VirtualMachineExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&computev1alpha1api20201201.VirtualMachine{},
		&v1alpha1api20201201storage.VirtualMachine{},
		&computev1beta20201201.VirtualMachine{},
		&v1beta20201201storage.VirtualMachine{}}
}