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

// Instance array details for the infrastructure requirement.
type AllOfInfrastructureRequirementDtoInstanceArray struct {
	// Label of the instance array.
	Label string `json:"label" yaml:"label"`
	// Minimum instance count for the instance array.
	MinInstanceCount float64 `json:"min_instance_count" yaml:"min_instance_count"`
	// Validation rules for instance count.
	InstanceCountValidation string `json:"instance_count_validation" yaml:"instance_count_validation"`
	// Instance count value.
	InstanceCount string `json:"instance_count" yaml:"instance_count"`
	// Minimum CPU required for the instance array.
	MinCpu float64 `json:"min_cpu" yaml:"min_cpu"`
	// Minimum RAM required for the instance array.
	MinRam float64 `json:"min_ram" yaml:"min_ram"`
	// Connected shared drive arrays.
	ConnectedSharedDriveArrays string `json:"connected_shared_drive_arrays,omitempty" yaml:"connected_shared_drive_arrays,omitempty"`
	// OS templates for the instance array.
	OsTemplates string `json:"os_templates" yaml:"os_templates"`
}
