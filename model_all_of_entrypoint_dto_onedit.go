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

// Actions for the onEdit event.
type AllOfEntrypointDtoOnedit struct {
	// Pre-deploy tasks.
	PreDeploy []TaskDto `json:"pre_deploy,omitempty" yaml:"pre_deploy,omitempty"`
	// Post-deploy tasks.
	PostDeploy []TaskDto `json:"post_deploy,omitempty" yaml:"post_deploy,omitempty"`
}
