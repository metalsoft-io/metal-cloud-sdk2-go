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

type UpdateExtensionInstanceDto struct {
	// The infrastructure ID.
	InfrastructureId float64 `json:"infrastructureId,omitempty" yaml:"infrastructureId,omitempty"`
	// The extension ID.
	ExtensionId float64 `json:"extensionId,omitempty" yaml:"extensionId,omitempty"`
	// The extension instance label. Will be automatically generated if not provided.
	Label string `json:"label,omitempty" yaml:"label,omitempty"`
	// Variables for extension instance.
	VariablesExtensionInstance *interface{} `json:"variablesExtensionInstance,omitempty" yaml:"variablesExtensionInstance,omitempty"`
	// Ansible bundles info.
	AnsibleBundlesInfo *interface{} `json:"ansibleBundlesInfo,omitempty" yaml:"ansibleBundlesInfo,omitempty"`
}
