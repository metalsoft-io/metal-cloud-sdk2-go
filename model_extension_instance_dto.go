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

type ExtensionInstanceDto struct {
	// The extension instance ID.
	Id float64 `json:"id" yaml:"id"`
	// The infrastructure ID.
	InfrastructureId float64 `json:"infrastructureId" yaml:"infrastructureId"`
	// The extension ID.
	ExtensionId float64 `json:"extensionId" yaml:"extensionId"`
	// The extension instance label. Will be automatically generated if not provided.
	Label string `json:"label" yaml:"label"`
	// Input variables values.
	InputVariables []ExtensionVariable `json:"inputVariables" yaml:"inputVariables"`
	// Output variables values.
	OutputVariables []ExtensionVariable `json:"outputVariables" yaml:"outputVariables"`
}
