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

type ExtensionTaskAnsible struct {
	// Asset for the Ansible task.
	Asset string `json:"asset,omitempty" yaml:"asset,omitempty"`
	// Playbook for the Ansible task.
	Playbook string `json:"playbook,omitempty" yaml:"playbook,omitempty"`
	// Execution Timeout.
	ExecutionTimeout int32 `json:"executionTimeout,omitempty" yaml:"executionTimeout,omitempty"`
	// Execution Timeout Tick.
	ExecutionTimeoutTick int32 `json:"executionTimeoutTick,omitempty" yaml:"executionTimeoutTick,omitempty"`
	// Version.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
