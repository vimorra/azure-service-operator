// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

type BlobContainer_STATUS_ARM struct {
	// Etag: Resource Etag.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the blob container.
	Properties *ContainerProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ContainerProperties_STATUS_ARM struct {
	// DefaultEncryptionScope: Default the container to use specified encryption scope for all writes.
	DefaultEncryptionScope *string `json:"defaultEncryptionScope,omitempty"`

	// Deleted: Indicates whether the blob container was deleted.
	Deleted *bool `json:"deleted,omitempty"`

	// DeletedTime: Blob container deletion time.
	DeletedTime *string `json:"deletedTime,omitempty"`

	// DenyEncryptionScopeOverride: Block override of encryption scope from the container default.
	DenyEncryptionScopeOverride *bool `json:"denyEncryptionScopeOverride,omitempty"`

	// HasImmutabilityPolicy: The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been
	// created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has
	// not been created for this container.
	HasImmutabilityPolicy *bool `json:"hasImmutabilityPolicy,omitempty"`

	// HasLegalHold: The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The
	// hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a
	// maximum of 1000 blob containers with hasLegalHold=true for a given account.
	HasLegalHold *bool `json:"hasLegalHold,omitempty"`

	// ImmutabilityPolicy: The ImmutabilityPolicy property of the container.
	ImmutabilityPolicy *ImmutabilityPolicyProperties_STATUS_ARM `json:"immutabilityPolicy,omitempty"`

	// ImmutableStorageWithVersioning: The object level immutability property of the container. The property is immutable and
	// can only be set to true at the container creation time. Existing containers must undergo a migration process.
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning_STATUS_ARM `json:"immutableStorageWithVersioning,omitempty"`

	// LastModifiedTime: Returns the date and time the container was last modified.
	LastModifiedTime *string `json:"lastModifiedTime,omitempty"`

	// LeaseDuration: Specifies whether the lease on a container is of infinite or fixed duration, only when the container is
	// leased.
	LeaseDuration *ContainerProperties_LeaseDuration_STATUS `json:"leaseDuration,omitempty"`

	// LeaseState: Lease state of the container.
	LeaseState *ContainerProperties_LeaseState_STATUS `json:"leaseState,omitempty"`

	// LeaseStatus: The lease status of the container.
	LeaseStatus *ContainerProperties_LeaseStatus_STATUS `json:"leaseStatus,omitempty"`

	// LegalHold: The LegalHold property of the container.
	LegalHold *LegalHoldProperties_STATUS_ARM `json:"legalHold,omitempty"`

	// Metadata: A name-value pair to associate with the container as metadata.
	Metadata map[string]string `json:"metadata,omitempty"`

	// PublicAccess: Specifies whether data in the container may be accessed publicly and the level of access.
	PublicAccess *ContainerProperties_PublicAccess_STATUS `json:"publicAccess,omitempty"`

	// RemainingRetentionDays: Remaining retention days for soft deleted blob container.
	RemainingRetentionDays *int `json:"remainingRetentionDays,omitempty"`

	// Version: The version of the deleted blob container.
	Version *string `json:"version,omitempty"`
}

type ContainerProperties_LeaseDuration_STATUS string

const (
	ContainerProperties_LeaseDuration_STATUS_Fixed    = ContainerProperties_LeaseDuration_STATUS("Fixed")
	ContainerProperties_LeaseDuration_STATUS_Infinite = ContainerProperties_LeaseDuration_STATUS("Infinite")
)

type ContainerProperties_LeaseState_STATUS string

const (
	ContainerProperties_LeaseState_STATUS_Available = ContainerProperties_LeaseState_STATUS("Available")
	ContainerProperties_LeaseState_STATUS_Breaking  = ContainerProperties_LeaseState_STATUS("Breaking")
	ContainerProperties_LeaseState_STATUS_Broken    = ContainerProperties_LeaseState_STATUS("Broken")
	ContainerProperties_LeaseState_STATUS_Expired   = ContainerProperties_LeaseState_STATUS("Expired")
	ContainerProperties_LeaseState_STATUS_Leased    = ContainerProperties_LeaseState_STATUS("Leased")
)

type ContainerProperties_LeaseStatus_STATUS string

const (
	ContainerProperties_LeaseStatus_STATUS_Locked   = ContainerProperties_LeaseStatus_STATUS("Locked")
	ContainerProperties_LeaseStatus_STATUS_Unlocked = ContainerProperties_LeaseStatus_STATUS("Unlocked")
)

type ContainerProperties_PublicAccess_STATUS string

const (
	ContainerProperties_PublicAccess_STATUS_Blob      = ContainerProperties_PublicAccess_STATUS("Blob")
	ContainerProperties_PublicAccess_STATUS_Container = ContainerProperties_PublicAccess_STATUS("Container")
	ContainerProperties_PublicAccess_STATUS_None      = ContainerProperties_PublicAccess_STATUS("None")
)

type ImmutabilityPolicyProperties_STATUS_ARM struct {
	// Etag: ImmutabilityPolicy Etag.
	Etag *string `json:"etag,omitempty"`

	// Properties: The properties of an ImmutabilityPolicy of a blob container.
	Properties *ImmutabilityPolicyProperty_STATUS_ARM `json:"properties,omitempty"`

	// UpdateHistory: The ImmutabilityPolicy update history of the blob container.
	UpdateHistory []UpdateHistoryProperty_STATUS_ARM `json:"updateHistory,omitempty"`
}

type ImmutableStorageWithVersioning_STATUS_ARM struct {
	// Enabled: This is an immutable property, when set to true it enables object level immutability at the container level.
	Enabled *bool `json:"enabled,omitempty"`

	// MigrationState: This property denotes the container level immutability to object level immutability migration state.
	MigrationState *ImmutableStorageWithVersioning_MigrationState_STATUS `json:"migrationState,omitempty"`

	// TimeStamp: Returns the date and time the object level immutability was enabled.
	TimeStamp *string `json:"timeStamp,omitempty"`
}

type LegalHoldProperties_STATUS_ARM struct {
	// HasLegalHold: The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The
	// hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a
	// maximum of 1000 blob containers with hasLegalHold=true for a given account.
	HasLegalHold *bool `json:"hasLegalHold,omitempty"`

	// Tags: The list of LegalHold tags of a blob container.
	Tags []TagProperty_STATUS_ARM `json:"tags,omitempty"`
}

type ImmutabilityPolicyProperty_STATUS_ARM struct {
	// AllowProtectedAppendWrites: This property can only be changed for unlocked time-based retention policies. When enabled,
	// new blocks can be written to an append blob while maintaining immutability protection and compliance. Only new blocks
	// can be added and any existing blocks cannot be modified or deleted. This property cannot be changed with
	// ExtendImmutabilityPolicy API
	AllowProtectedAppendWrites *bool `json:"allowProtectedAppendWrites,omitempty"`

	// ImmutabilityPeriodSinceCreationInDays: The immutability period for the blobs in the container since the policy creation,
	// in days.
	ImmutabilityPeriodSinceCreationInDays *int `json:"immutabilityPeriodSinceCreationInDays,omitempty"`

	// State: The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
	State *ImmutabilityPolicyProperty_State_STATUS `json:"state,omitempty"`
}

type ImmutableStorageWithVersioning_MigrationState_STATUS string

const (
	ImmutableStorageWithVersioning_MigrationState_STATUS_Completed  = ImmutableStorageWithVersioning_MigrationState_STATUS("Completed")
	ImmutableStorageWithVersioning_MigrationState_STATUS_InProgress = ImmutableStorageWithVersioning_MigrationState_STATUS("InProgress")
)

type TagProperty_STATUS_ARM struct {
	// ObjectIdentifier: Returns the Object ID of the user who added the tag.
	ObjectIdentifier *string `json:"objectIdentifier,omitempty"`

	// Tag: The tag value.
	Tag *string `json:"tag,omitempty"`

	// TenantId: Returns the Tenant ID that issued the token for the user who added the tag.
	TenantId *string `json:"tenantId,omitempty"`

	// Timestamp: Returns the date and time the tag was added.
	Timestamp *string `json:"timestamp,omitempty"`

	// Upn: Returns the User Principal Name of the user who added the tag.
	Upn *string `json:"upn,omitempty"`
}

type UpdateHistoryProperty_STATUS_ARM struct {
	// ImmutabilityPeriodSinceCreationInDays: The immutability period for the blobs in the container since the policy creation,
	// in days.
	ImmutabilityPeriodSinceCreationInDays *int `json:"immutabilityPeriodSinceCreationInDays,omitempty"`

	// ObjectIdentifier: Returns the Object ID of the user who updated the ImmutabilityPolicy.
	ObjectIdentifier *string `json:"objectIdentifier,omitempty"`

	// TenantId: Returns the Tenant ID that issued the token for the user who updated the ImmutabilityPolicy.
	TenantId *string `json:"tenantId,omitempty"`

	// Timestamp: Returns the date and time the ImmutabilityPolicy was updated.
	Timestamp *string `json:"timestamp,omitempty"`

	// Update: The ImmutabilityPolicy update type of a blob container, possible values include: put, lock and extend.
	Update *UpdateHistoryProperty_Update_STATUS `json:"update,omitempty"`

	// Upn: Returns the User Principal Name of the user who updated the ImmutabilityPolicy.
	Upn *string `json:"upn,omitempty"`
}

type UpdateHistoryProperty_Update_STATUS string

const (
	UpdateHistoryProperty_Update_STATUS_Extend = UpdateHistoryProperty_Update_STATUS("extend")
	UpdateHistoryProperty_Update_STATUS_Lock   = UpdateHistoryProperty_Update_STATUS("lock")
	UpdateHistoryProperty_Update_STATUS_Put    = UpdateHistoryProperty_Update_STATUS("put")
)