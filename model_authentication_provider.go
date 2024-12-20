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

type AuthenticationProvider struct {
	// Authentication provider name
	Name string `json:"name" yaml:"name"`
	// True if the provider is enabled
	Enabled bool `json:"enabled" yaml:"enabled"`
	// Permitted domains
	Domains []string `json:"domains" yaml:"domains"`
}
