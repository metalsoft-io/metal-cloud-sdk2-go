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

type ExtensionSharedDrive struct {
	// Label of the shared drive array.
	Label string `json:"label" yaml:"label"`
	// Size of the shared drive array in GB..
	SizeGb float64 `json:"sizeGb" yaml:"sizeGb"`
}