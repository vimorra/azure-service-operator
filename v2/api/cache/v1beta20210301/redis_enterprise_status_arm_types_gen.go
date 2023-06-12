// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210301

// Deprecated version of RedisEnterprise_STATUS. Use v1api20210301.RedisEnterprise_STATUS instead
type RedisEnterprise_STATUS_ARM struct {
	Id         *string                       `json:"id,omitempty"`
	Location   *string                       `json:"location,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	Properties *ClusterProperties_STATUS_ARM `json:"properties,omitempty"`
	Sku        *Sku_STATUS_ARM               `json:"sku,omitempty"`
	Tags       map[string]string             `json:"tags,omitempty"`
	Type       *string                       `json:"type,omitempty"`
	Zones      []string                      `json:"zones,omitempty"`
}

// Deprecated version of ClusterProperties_STATUS. Use v1api20210301.ClusterProperties_STATUS instead
type ClusterProperties_STATUS_ARM struct {
	HostName                   *string                                     `json:"hostName,omitempty"`
	MinimumTlsVersion          *ClusterProperties_MinimumTlsVersion_STATUS `json:"minimumTlsVersion,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_ARM      `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *ProvisioningState_STATUS                   `json:"provisioningState,omitempty"`
	RedisVersion               *string                                     `json:"redisVersion,omitempty"`
	ResourceState              *ResourceState_STATUS                       `json:"resourceState,omitempty"`
}

// Deprecated version of Sku_STATUS. Use v1api20210301.Sku_STATUS instead
type Sku_STATUS_ARM struct {
	Capacity *int             `json:"capacity,omitempty"`
	Name     *Sku_Name_STATUS `json:"name,omitempty"`
}

// Deprecated version of PrivateEndpointConnection_STATUS. Use v1api20210301.PrivateEndpointConnection_STATUS instead
type PrivateEndpointConnection_STATUS_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of Sku_Name_STATUS. Use v1api20210301.Sku_Name_STATUS instead
type Sku_Name_STATUS string

const (
	Sku_Name_STATUS_EnterpriseFlash_F1500 = Sku_Name_STATUS("EnterpriseFlash_F1500")
	Sku_Name_STATUS_EnterpriseFlash_F300  = Sku_Name_STATUS("EnterpriseFlash_F300")
	Sku_Name_STATUS_EnterpriseFlash_F700  = Sku_Name_STATUS("EnterpriseFlash_F700")
	Sku_Name_STATUS_Enterprise_E10        = Sku_Name_STATUS("Enterprise_E10")
	Sku_Name_STATUS_Enterprise_E100       = Sku_Name_STATUS("Enterprise_E100")
	Sku_Name_STATUS_Enterprise_E20        = Sku_Name_STATUS("Enterprise_E20")
	Sku_Name_STATUS_Enterprise_E50        = Sku_Name_STATUS("Enterprise_E50")
)