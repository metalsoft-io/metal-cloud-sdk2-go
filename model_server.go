/*
 * MetalSoft REST API
 *
 * MetalSoft REST API documentation
 *
 * API version: 2.0
 * Contact: support@metalsoft.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sdk2

type Server struct {
	ServerId float64 `json:"serverId" yaml:"serverId"`
	ServerTypeId float64 `json:"serverTypeId" yaml:"serverTypeId"`
	DatacenterName string `json:"datacenterName" yaml:"datacenterName"`
	ServerUUID string `json:"serverUUID" yaml:"serverUUID"`
	SerialNumber *string `json:"serialNumber,omitempty" yaml:"serialNumber,omitempty"`
	ManagementAddress string `json:"managementAddress" yaml:"managementAddress"`
	Username string `json:"username" yaml:"username"`
	PasswordEncrypted string `json:"passwordEncrypted" yaml:"passwordEncrypted"`
	ImpiVersion *string `json:"impiVersion,omitempty" yaml:"impiVersion,omitempty"`
	RamGbytes float64 `json:"ramGbytes" yaml:"ramGbytes"`
	ProcessorCount float64 `json:"processorCount" yaml:"processorCount"`
	ProcessorCoreMhz float64 `json:"processorCoreMhz" yaml:"processorCoreMhz"`
	ProcessorCoreCount float64 `json:"processorCoreCount" yaml:"processorCoreCount"`
	ProcessorName *string `json:"processorName,omitempty" yaml:"processorName,omitempty"`
	ProcessorCpuMark *float64 `json:"processorCpuMark,omitempty" yaml:"processorCpuMark,omitempty"`
	ProcessorThreads *float64 `json:"processorThreads,omitempty" yaml:"processorThreads,omitempty"`
	DiskCount float64 `json:"diskCount" yaml:"diskCount"`
	MgmtSnmpPort *float64 `json:"mgmtSnmpPort,omitempty" yaml:"mgmtSnmpPort,omitempty"`
	MgmtSnmpPasswordEncrypted *string `json:"mgmtSnmpPasswordEncrypted,omitempty" yaml:"mgmtSnmpPasswordEncrypted,omitempty"`
	BmcMacAddress string `json:"bmcMacAddress,omitempty" yaml:"bmcMacAddress,omitempty"`
	ServerMetricsMetadata *interface{} `json:"serverMetricsMetadata" yaml:"serverMetricsMetadata"`
	Vendor *string `json:"vendor,omitempty" yaml:"vendor,omitempty"`
	VendorSkuId *string `json:"vendorSkuId,omitempty" yaml:"vendorSkuId,omitempty"`
	Model *string `json:"model,omitempty" yaml:"model,omitempty"`
	VncPort *float64 `json:"vncPort,omitempty" yaml:"vncPort,omitempty"`
	VncPasswordEncrypted *float64 `json:"vncPasswordEncrypted,omitempty" yaml:"vncPasswordEncrypted,omitempty"`
	IsBasicCampusEndpoint float64 `json:"isBasicCampusEndpoint" yaml:"isBasicCampusEndpoint"`
	ServerCleanupPolicyId *float64 `json:"serverCleanupPolicyId,omitempty" yaml:"serverCleanupPolicyId,omitempty"`
	RequiresReRegister float64 `json:"requiresReRegister" yaml:"requiresReRegister"`
	ServerSupportsSol float64 `json:"serverSupportsSol" yaml:"serverSupportsSol"`
	ServerSupportsVirtualMedia float64 `json:"serverSupportsVirtualMedia" yaml:"serverSupportsVirtualMedia"`
	ServerSupportsOobProvisioning float64 `json:"serverSupportsOobProvisioning" yaml:"serverSupportsOobProvisioning"`
	BiosInfo *interface{} `json:"biosInfo,omitempty" yaml:"biosInfo,omitempty"`
	VendorInfo *interface{} `json:"vendorInfo,omitempty" yaml:"vendorInfo,omitempty"`
	ServerClass string `json:"serverClass" yaml:"serverClass"`
	ServerStatus string `json:"serverStatus" yaml:"serverStatus"`
	ServerComments *string `json:"serverComments,omitempty" yaml:"serverComments,omitempty"`
	ServerCapacityMbps *float64 `json:"serverCapacityMbps,omitempty" yaml:"serverCapacityMbps,omitempty"`
	ServerDiskWipe float64 `json:"serverDiskWipe" yaml:"serverDiskWipe"`
	ServerDiskCount *float64 `json:"serverDiskCount,omitempty" yaml:"serverDiskCount,omitempty"`
	AdministrationState string `json:"administrationState" yaml:"administrationState"`
	ServerDhcpStatus string `json:"serverDhcpStatus" yaml:"serverDhcpStatus"`
	GpuCount float64 `json:"gpuCount" yaml:"gpuCount"`
	GpuInfo []interface{} `json:"gpuInfo" yaml:"gpuInfo"`
	Submodel *string `json:"submodel,omitempty" yaml:"submodel,omitempty"`
	SupportsFcProvisioning float64 `json:"supportsFcProvisioning" yaml:"supportsFcProvisioning"`
	BootLastUpdateTimestamp *string `json:"bootLastUpdateTimestamp,omitempty" yaml:"bootLastUpdateTimestamp,omitempty"`
	RegisteredTimestamp *string `json:"registeredTimestamp,omitempty" yaml:"registeredTimestamp,omitempty"`
	ServerCreatedTimestamp string `json:"serverCreatedTimestamp" yaml:"serverCreatedTimestamp"`
	PowerStatus string `json:"powerStatus" yaml:"powerStatus"`
	PowerStatusLastUpdateTimestamp string `json:"powerStatusLastUpdateTimestamp" yaml:"powerStatusLastUpdateTimestamp"`
}
