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

type ExtensionDefinitionDto struct {
	// The kind of extension.
	Kind string `json:"kind" yaml:"kind"`
	// Schema version of the extension.
	SchemaVersion string `json:"schemaVersion" yaml:"schemaVersion"`
	// Name of the extension.
	Name string `json:"name" yaml:"name"`
	// Label of the extension.
	Label string `json:"label,omitempty" yaml:"label,omitempty"`
	// Type of the extension.
	ExtensionType string `json:"extensionType" yaml:"extensionType"`
	// Vendor of the extension.
	Vendor string `json:"vendor" yaml:"vendor"`
	// Version of the extension.
	ExtensionVersion string `json:"extensionVersion" yaml:"extensionVersion"`
	// Description of the extension.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	// Icon of the extension.
	Icon string `json:"icon" yaml:"icon"`
	Dependencies *ExtensionDependency `json:"dependencies" yaml:"dependencies"`
	// List of inputs for the platform service.
	Inputs []ExtensionInput `json:"inputs" yaml:"inputs"`
	// List of outputs for the platform service.
	Outputs []ExtensionOutput `json:"outputs" yaml:"outputs"`
	Infrastructure *ExtensionInfrastructure `json:"infrastructure" yaml:"infrastructure"`
	// List of assets for the platform service.
	Assets []ExtensionAsset `json:"assets" yaml:"assets"`
	OnCreate *ExtensionActions `json:"onCreate,omitempty" yaml:"onCreate,omitempty"`
	OnEdit *ExtensionActions `json:"onEdit,omitempty" yaml:"onEdit,omitempty"`
	OnDelete *ExtensionActions `json:"onDelete,omitempty" yaml:"onDelete,omitempty"`
}