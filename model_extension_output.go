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

type ExtensionOutput struct {
	// Label of the output.
	Label string `json:"label" yaml:"label"`
	// Name of the output.
	Name string `json:"name" yaml:"name"`
	// Type of the output.
	OutputType string `json:"outputType" yaml:"outputType"`
	// Template of the output.
	Template string `json:"template,omitempty" yaml:"template,omitempty"`
}
