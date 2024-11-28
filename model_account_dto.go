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

type AccountDto struct {
	// Account ID
	Id string `json:"id" yaml:"id"`
	// The ID of the parent account
	ParentAccountId float64 `json:"parentAccountId,omitempty" yaml:"parentAccountId,omitempty"`
	// The name of the account
	Name string `json:"name" yaml:"name"`
	// The code of the account
	Code string `json:"code,omitempty" yaml:"code,omitempty"`
	// The fiscal number of the account
	FiscalNumber string `json:"fiscalNumber,omitempty" yaml:"fiscalNumber,omitempty"`
	Address *AccountAddressDto `json:"address,omitempty" yaml:"address,omitempty"`
	// The user ID of the primary contact
	PrimaryContactId float64 `json:"primaryContactId,omitempty" yaml:"primaryContactId,omitempty"`
	// The user ID of the secondary contact
	SecondaryContactId float64 `json:"secondaryContactId,omitempty" yaml:"secondaryContactId,omitempty"`
	// Whether the account is archived
	Archived float64 `json:"archived,omitempty" yaml:"archived,omitempty"`
	Limits *AccountLimits `json:"limits" yaml:"limits"`
}
