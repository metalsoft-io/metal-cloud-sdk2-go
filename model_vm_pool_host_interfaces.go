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

type VmPoolHostInterfaces struct {
	// VM Pool Host Interface ID
	Id float64 `json:"id" yaml:"id"`
	// VM Pool Host ID
	HostId float64 `json:"hostId" yaml:"hostId"`
	// Name of the VM Pool Host Interface
	Name string `json:"name" yaml:"name"`
	// MAC Address of the VM Pool Host Interface
	MacAddress string `json:"macAddress" yaml:"macAddress"`
	// Fabric of the VM Pool Host Interface
	Fabric string `json:"fabric" yaml:"fabric"`
	Links *interface{} `json:"links" yaml:"links"`
}
