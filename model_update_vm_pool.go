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

type UpdateVmPool struct {
	// Datacenter of the VM Pool
	Description *string `json:"description,omitempty" yaml:"description,omitempty"`
	// Host of the VM Pool
	ManagementHost string `json:"managementHost,omitempty" yaml:"managementHost,omitempty"`
	// Port of the VM Pool
	ManagementPort float64 `json:"managementPort,omitempty" yaml:"managementPort,omitempty"`
	// Certificate of the VM Pool
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	// Private key of the VM Pool
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`
	// Status of the VM Pool
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
	// Flag to indicate if the VM Pool is in maintenance mode. 1 for true, 0 for false. Default is 0.
	InMaintenance float64 `json:"inMaintenance,omitempty" yaml:"inMaintenance,omitempty"`
	// Flag to indicate if the VM Pool is experimental. 1 for true, 0 for false. Default is 0.
	IsExperimental float64 `json:"isExperimental,omitempty" yaml:"isExperimental,omitempty"`
	// Tags for the VM Pool.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
