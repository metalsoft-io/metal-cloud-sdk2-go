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

type InfrastructureDto struct {
	// Infrastructure Id
	InfrastructureId float64 `json:"infrastructureId" yaml:"infrastructureId"`
	// Infrastructure Change ID
	InfrastructureChangeId float64 `json:"infrastructureChangeId" yaml:"infrastructureChangeId"`
	// Name of the datacenter
	DatacenterName string `json:"datacenterName" yaml:"datacenterName"`
	// Owner of the infrastructure
	UserIdOwner *float64 `json:"userIdOwner,omitempty" yaml:"userIdOwner,omitempty"`
	// Status of the infrastructure
	InfrastructureServiceStatus string `json:"infrastructureServiceStatus" yaml:"infrastructureServiceStatus"`
	// Operation object of the infrastructure
	Operation *interface{} `json:"operation" yaml:"operation"`
}
