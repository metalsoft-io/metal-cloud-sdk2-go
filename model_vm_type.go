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

type VmType struct {
	// VM Type ID
	Id float64 `json:"id"`
	// Name of the VM Pool type
	Name string `json:"name"`
	// Display name of the VM Pool type
	DisplayName string `json:"displayName"`
	// Label of the VM Pool type
	Label string `json:"label"`
	// Number of CPU cores for the VM Pool type
	CpuCores float64 `json:"cpuCores"`
	// RAM in GB for the VM Pool type
	RamGB float64 `json:"ramGB"`
	// Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0.
	IsExperimental float64 `json:"isExperimental"`
	// Tags for the VM Type.
	Tags []string `json:"tags"`
	// Flag to indicate if the VM Pool is for unmanaged VMs only. 1 for true, 0 for false. Default is 0.
	ForUnmanagedVMsOnly float64 `json:"forUnmanagedVMsOnly"`
	Links *interface{} `json:"links"`
}
