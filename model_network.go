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

type Network struct {
	// Id of the Network.
	NetworkId float64 `json:"networkId" yaml:"networkId"`
	// Type of the Network.
	NetworkType string `json:"networkType" yaml:"networkType"`
	// Id of the Infrastructure.
	InfrastructureId float64 `json:"infrastructureId" yaml:"infrastructureId"`
	// Network Service status.
	ServiceStatus string `json:"serviceStatus" yaml:"serviceStatus"`
	// Name of the Network.
	Label string `json:"label" yaml:"label"`
	// Subdomain of the Network.
	Subdomain string `json:"subdomain" yaml:"subdomain"`
	// Permanent subdomain of the Network.
	SubdomainPermanent string `json:"subdomainPermanent" yaml:"subdomainPermanent"`
	// Timestamp of the Network creation.
	CreatedTimestamp string `json:"createdTimestamp" yaml:"createdTimestamp"`
	// Timestamp of the Network update.
	UpdatedTimestamp string `json:"updatedTimestamp" yaml:"updatedTimestamp"`
}
