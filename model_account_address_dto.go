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

type AccountAddressDto struct {
	// The first line of the account address
	AddressLine1 string `json:"addressLine1,omitempty"`
	// The second line of the account address
	AddressLine2 string `json:"addressLine2,omitempty"`
	// The postal code of the account address
	PostalCode string `json:"postalCode,omitempty"`
	// The state of the account address
	State string `json:"state,omitempty"`
	// The country of the account address
	Country string `json:"country,omitempty"`
}
