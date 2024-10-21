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

// Run a webhook task.
type AllOfTaskDtoRunWebhook struct {
	// Asset for the webhook task.
	Asset string `json:"asset" yaml:"asset"`
	// Arguments for the webhook task.
	Args string `json:"args" yaml:"args"`
}
