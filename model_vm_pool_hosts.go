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

type VmPoolHosts struct {
	// VM Pool Host ID
	Id float64 `json:"id"`
	// Name of the VM Pool Host
	Name string `json:"name"`
	// Address of the VM Pool Host
	Address string `json:"address"`
	// Port of the VM Pool Host
	Port float64 `json:"port"`
	// VM Pool ID
	PoolId float64 `json:"poolId"`
	// Roles of the VM Pool Host
	Roles []string `json:"roles"`
	// Failure domain of the VM Pool Host
	FailureDomain string `json:"failureDomain"`
	// Architecture of the VM Pool Host
	Architecture string `json:"architecture"`
	// Flag specifying if the VM Pool Host is database
	Database float64 `json:"database"`
	// Status of the VM Pool Host
	Status string `json:"status"`
	// Description of the VM Pool Host
	Description string `json:"description,omitempty"`
	// Total RAM GB of the VM Pool Host
	TotalRamGB float64 `json:"totalRamGB"`
	// Free RAM GB of the VM Pool Host
	FreeRamGB float64 `json:"freeRamGB"`
	// Used RAM GB of the VM Pool Host
	UsedRamGB float64 `json:"usedRamGB"`
	// Total Space GB of the VM Pool Host
	TotalSpaceGB float64 `json:"totalSpaceGB"`
	// Used Space GB of the VM Pool Host
	UsedSpaceGB float64 `json:"usedSpaceGB"`
	// Free Space GB of the VM Pool Host
	FreeSpaceGB float64 `json:"freeSpaceGB"`
	// Timestamp when the VM Pool Host was updated
	UpdatedTimestamp string `json:"updatedTimestamp"`
}