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

type VmPoolHostInterfacesDto struct {
	// VM Pool Host Interface ID
	Id float64 `json:"id"`
	// VM Pool Host ID
	HostId float64 `json:"hostId"`
	// Name of the VM Pool Host Interface
	Name string `json:"name"`
	// MAC Address of the VM Pool Host Interface
	MacAddress string `json:"macAddress"`
	// Fabric of the VM Pool Host Interface
	Fabric string `json:"fabric"`
}
