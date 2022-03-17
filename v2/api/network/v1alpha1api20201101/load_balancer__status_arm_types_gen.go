// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

type LoadBalancer_StatusARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Location: Resource location.
	Location *string `json:"location,omitempty"`

	//Name: Resource name.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of load balancer.
	Properties *LoadBalancerPropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Sku: The load balancer SKU.
	Sku *LoadBalancerSku_StatusARM `json:"sku,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type.
	Type *string `json:"type,omitempty"`
}

type ExtendedLocation_StatusARM struct {
	//Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	//Type: The type of the extended location.
	Type *ExtendedLocationType_Status `json:"type,omitempty"`
}

type LoadBalancerPropertiesFormat_StatusARM struct {
	//BackendAddressPools: Collection of backend address pools used by a load balancer.
	BackendAddressPools []BackendAddressPool_Status_LoadBalancer_SubResourceEmbeddedARM `json:"backendAddressPools,omitempty"`

	//FrontendIPConfigurations: Object representing the frontend IPs to be used for the load balancer.
	FrontendIPConfigurations []FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbeddedARM `json:"frontendIPConfigurations,omitempty"`

	//InboundNatPools: Defines an external port range for inbound NAT to a single backend port on NICs associated with a load
	//balancer. Inbound NAT rules are created automatically for each NIC associated with the Load Balancer using an external
	//port from this range. Defining an Inbound NAT pool on your Load Balancer is mutually exclusive with defining inbound Nat
	//rules. Inbound NAT pools are referenced from virtual machine scale sets. NICs that are associated with individual
	//virtual machines cannot reference an inbound NAT pool. They have to reference individual inbound NAT rules.
	InboundNatPools []InboundNatPool_StatusARM `json:"inboundNatPools,omitempty"`

	//InboundNatRules: Collection of inbound NAT Rules used by a load balancer. Defining inbound NAT rules on your load
	//balancer is mutually exclusive with defining an inbound NAT pool. Inbound NAT pools are referenced from virtual machine
	//scale sets. NICs that are associated with individual virtual machines cannot reference an Inbound NAT pool. They have to
	//reference individual inbound NAT rules.
	InboundNatRules []InboundNatRule_Status_LoadBalancer_SubResourceEmbeddedARM `json:"inboundNatRules,omitempty"`

	//LoadBalancingRules: Object collection representing the load balancing rules Gets the provisioning.
	LoadBalancingRules []LoadBalancingRule_StatusARM `json:"loadBalancingRules,omitempty"`

	//OutboundRules: The outbound rules.
	OutboundRules []OutboundRule_StatusARM `json:"outboundRules,omitempty"`

	//Probes: Collection of probe objects used in the load balancer.
	Probes []Probe_StatusARM `json:"probes,omitempty"`

	//ProvisioningState: The provisioning state of the load balancer resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	//ResourceGuid: The resource GUID property of the load balancer resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`
}

type LoadBalancerSku_StatusARM struct {
	//Name: Name of a load balancer SKU.
	Name *LoadBalancerSkuStatusName `json:"name,omitempty"`

	//Tier: Tier of a load balancer SKU.
	Tier *LoadBalancerSkuStatusTier `json:"tier,omitempty"`
}

type BackendAddressPool_Status_LoadBalancer_SubResourceEmbeddedARM struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type ExtendedLocationType_Status string

const ExtendedLocationType_StatusEdgeZone = ExtendedLocationType_Status("EdgeZone")

type FrontendIPConfiguration_Status_LoadBalancer_SubResourceEmbeddedARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within the set of frontend IP configurations used by the load balancer.
	//This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of the load balancer probe.
	Properties *FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM `json:"properties,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`

	//Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type InboundNatPool_StatusARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within the set of inbound NAT pools used by the load balancer. This name
	//can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of load balancer inbound nat pool.
	Properties *InboundNatPoolPropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type InboundNatRule_Status_LoadBalancer_SubResourceEmbeddedARM struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type LoadBalancerSkuStatusName string

const (
	LoadBalancerSkuStatusNameBasic    = LoadBalancerSkuStatusName("Basic")
	LoadBalancerSkuStatusNameStandard = LoadBalancerSkuStatusName("Standard")
)

type LoadBalancerSkuStatusTier string

const (
	LoadBalancerSkuStatusTierGlobal   = LoadBalancerSkuStatusTier("Global")
	LoadBalancerSkuStatusTierRegional = LoadBalancerSkuStatusTier("Regional")
)

type LoadBalancingRule_StatusARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within the set of load balancing rules used by the load balancer. This
	//name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of load balancer load balancing rule.
	Properties *LoadBalancingRulePropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type OutboundRule_StatusARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within the set of outbound rules used by the load balancer. This name can
	//be used to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of load balancer outbound rule.
	Properties *OutboundRulePropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type Probe_StatusARM struct {
	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within the set of probes used by the load balancer. This name can be used
	//to access the resource.
	Name *string `json:"name,omitempty"`

	//Properties: Properties of load balancer probe.
	Properties *ProbePropertiesFormat_StatusARM `json:"properties,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type FrontendIPConfigurationPropertiesFormat_Status_LoadBalancer_SubResourceEmbeddedARM struct {
	//InboundNatPools: An array of references to inbound pools that use this frontend IP.
	InboundNatPools []SubResource_StatusARM `json:"inboundNatPools,omitempty"`

	//InboundNatRules: An array of references to inbound rules that use this frontend IP.
	InboundNatRules []SubResource_StatusARM `json:"inboundNatRules,omitempty"`

	//LoadBalancingRules: An array of references to load balancing rules that use this frontend IP.
	LoadBalancingRules []SubResource_StatusARM `json:"loadBalancingRules,omitempty"`

	//OutboundRules: An array of references to outbound rules that use this frontend IP.
	OutboundRules []SubResource_StatusARM `json:"outboundRules,omitempty"`

	//PrivateIPAddress: The private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	//PrivateIPAddressVersion: Whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.
	PrivateIPAddressVersion *IPVersion_Status `json:"privateIPAddressVersion,omitempty"`

	//PrivateIPAllocationMethod: The Private IP allocation method.
	PrivateIPAllocationMethod *IPAllocationMethod_Status `json:"privateIPAllocationMethod,omitempty"`

	//ProvisioningState: The provisioning state of the frontend IP configuration resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	//PublicIPAddress: The reference to the Public IP resource.
	PublicIPAddress *PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM `json:"publicIPAddress,omitempty"`

	//PublicIPPrefix: The reference to the Public IP Prefix resource.
	PublicIPPrefix *SubResource_StatusARM `json:"publicIPPrefix,omitempty"`

	//Subnet: The reference to the subnet resource.
	Subnet *Subnet_Status_LoadBalancer_SubResourceEmbeddedARM `json:"subnet,omitempty"`
}

type InboundNatPoolPropertiesFormat_StatusARM struct {
	//BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 1 and 65535.
	BackendPort *int `json:"backendPort,omitempty"`

	//EnableFloatingIP: Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL
	//AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
	//This setting can't be changed after you create the endpoint.
	EnableFloatingIP *bool `json:"enableFloatingIP,omitempty"`

	//EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
	//element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `json:"enableTcpReset,omitempty"`

	//FrontendIPConfiguration: A reference to frontend IP addresses.
	FrontendIPConfiguration *SubResource_StatusARM `json:"frontendIPConfiguration,omitempty"`

	//FrontendPortRangeEnd: The last port number in the range of external ports that will be used to provide Inbound Nat to
	//NICs associated with a load balancer. Acceptable values range between 1 and 65535.
	FrontendPortRangeEnd *int `json:"frontendPortRangeEnd,omitempty"`

	//FrontendPortRangeStart: The first port number in the range of external ports that will be used to provide Inbound Nat to
	//NICs associated with a load balancer. Acceptable values range between 1 and 65534.
	FrontendPortRangeStart *int `json:"frontendPortRangeStart,omitempty"`

	//IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
	//default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	//Protocol: The reference to the transport protocol used by the inbound NAT pool.
	Protocol *TransportProtocol_Status `json:"protocol,omitempty"`

	//ProvisioningState: The provisioning state of the inbound NAT pool resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`
}

type LoadBalancingRulePropertiesFormat_StatusARM struct {
	//BackendAddressPool: A reference to a pool of DIPs. Inbound traffic is randomly load balanced across IPs in the backend
	//IPs.
	BackendAddressPool *SubResource_StatusARM `json:"backendAddressPool,omitempty"`

	//BackendPort: The port used for internal connections on the endpoint. Acceptable values are between 0 and 65535. Note
	//that value 0 enables "Any Port".
	BackendPort *int `json:"backendPort,omitempty"`

	//DisableOutboundSnat: Configures SNAT for the VMs in the backend pool to use the publicIP address specified in the
	//frontend of the load balancing rule.
	DisableOutboundSnat *bool `json:"disableOutboundSnat,omitempty"`

	//EnableFloatingIP: Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL
	//AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server.
	//This setting can't be changed after you create the endpoint.
	EnableFloatingIP *bool `json:"enableFloatingIP,omitempty"`

	//EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
	//element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `json:"enableTcpReset,omitempty"`

	//FrontendIPConfiguration: A reference to frontend IP addresses.
	FrontendIPConfiguration *SubResource_StatusARM `json:"frontendIPConfiguration,omitempty"`

	//FrontendPort: The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer.
	//Acceptable values are between 0 and 65534. Note that value 0 enables "Any Port".
	FrontendPort *int `json:"frontendPort,omitempty"`

	//IdleTimeoutInMinutes: The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The
	//default value is 4 minutes. This element is only used when the protocol is set to TCP.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	//LoadDistribution: The load distribution policy for this rule.
	LoadDistribution *LoadBalancingRulePropertiesFormatStatusLoadDistribution `json:"loadDistribution,omitempty"`

	//Probe: The reference to the load balancer probe used by the load balancing rule.
	Probe *SubResource_StatusARM `json:"probe,omitempty"`

	//Protocol: The reference to the transport protocol used by the load balancing rule.
	Protocol *TransportProtocol_Status `json:"protocol,omitempty"`

	//ProvisioningState: The provisioning state of the load balancing rule resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`
}

type OutboundRulePropertiesFormat_StatusARM struct {
	//AllocatedOutboundPorts: The number of outbound ports to be used for NAT.
	AllocatedOutboundPorts *int `json:"allocatedOutboundPorts,omitempty"`

	//BackendAddressPool: A reference to a pool of DIPs. Outbound traffic is randomly load balanced across IPs in the backend
	//IPs.
	BackendAddressPool *SubResource_StatusARM `json:"backendAddressPool,omitempty"`

	//EnableTcpReset: Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This
	//element is only used when the protocol is set to TCP.
	EnableTcpReset *bool `json:"enableTcpReset,omitempty"`

	//FrontendIPConfigurations: The Frontend IP addresses of the load balancer.
	FrontendIPConfigurations []SubResource_StatusARM `json:"frontendIPConfigurations,omitempty"`

	//IdleTimeoutInMinutes: The timeout for the TCP idle connection.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	//Protocol: The protocol for the outbound rule in load balancer.
	Protocol *OutboundRulePropertiesFormatStatusProtocol `json:"protocol,omitempty"`

	//ProvisioningState: The provisioning state of the outbound rule resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`
}

type ProbePropertiesFormat_StatusARM struct {
	//IntervalInSeconds: The interval, in seconds, for how frequently to probe the endpoint for health status. Typically, the
	//interval is slightly less than half the allocated timeout period (in seconds) which allows two full probes before taking
	//the instance out of rotation. The default value is 15, the minimum value is 5.
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`

	//LoadBalancingRules: The load balancer rules that use this probe.
	LoadBalancingRules []SubResource_StatusARM `json:"loadBalancingRules,omitempty"`

	//NumberOfProbes: The number of probes where if no response, will result in stopping further traffic from being delivered
	//to the endpoint. This values allows endpoints to be taken out of rotation faster or slower than the typical times used
	//in Azure.
	NumberOfProbes *int `json:"numberOfProbes,omitempty"`

	//Port: The port for communicating the probe. Possible values range from 1 to 65535, inclusive.
	Port *int `json:"port,omitempty"`

	//Protocol: The protocol of the end point. If 'Tcp' is specified, a received ACK is required for the probe to be
	//successful. If 'Http' or 'Https' is specified, a 200 OK response from the specifies URI is required for the probe to be
	//successful.
	Protocol *ProbePropertiesFormatStatusProtocol `json:"protocol,omitempty"`

	//ProvisioningState: The provisioning state of the probe resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	//RequestPath: The URI used for requesting health status from the VM. Path is required if a protocol is set to http.
	//Otherwise, it is not allowed. There is no default value.
	RequestPath *string `json:"requestPath,omitempty"`
}

type PublicIPAddress_Status_LoadBalancer_SubResourceEmbeddedARM struct {
	//ExtendedLocation: The extended location of the public ip address.
	ExtendedLocation *ExtendedLocation_StatusARM `json:"extendedLocation,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Sku: The public IP address SKU.
	Sku *PublicIPAddressSku_StatusARM `json:"sku,omitempty"`

	//Zones: A list of availability zones denoting the IP allocated for the resource needs to come from.
	Zones []string `json:"zones,omitempty"`
}

type Subnet_Status_LoadBalancer_SubResourceEmbeddedARM struct {
	//Id: Resource ID.
	Id *string `json:"id,omitempty"`
}