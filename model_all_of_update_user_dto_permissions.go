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

// The new permissions of the user.
type AllOfUpdateUserDtoPermissions struct {
	// The new special permissions of the user.
	SpecialPermissions []UserResourcePermissionDto `json:"specialPermissions,omitempty" yaml:"specialPermissions,omitempty"`
	// The new password reveal permissions of the user.
	PasswordRevealPermissions []UserResourcePermissionDto `json:"passwordRevealPermissions,omitempty" yaml:"passwordRevealPermissions,omitempty"`
}
