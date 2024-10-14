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

type CreateVmPool struct {
	// Id of the site for the VM Pool
	SiteId float64 `json:"siteId,omitempty"`
	// Host of the VM Pool
	ManagementHost string `json:"managementHost,omitempty"`
	// Port of the VM Pool
	ManagementPort float64 `json:"managementPort,omitempty"`
	// Name of the VM Pool
	Name string `json:"name,omitempty"`
	// Description of the VM Pool
	Description string `json:"description,omitempty"`
	// Type of the VM Pool
	Type_ string `json:"type,omitempty"`
	// Certificate of the VM Pool
	Certificate string `json:"certificate,omitempty"`
	// Private key of the VM Pool
	PrivateKey string `json:"privateKey,omitempty"`
}
