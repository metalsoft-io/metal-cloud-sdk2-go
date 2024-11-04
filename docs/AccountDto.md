# AccountDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float64** | Account ID | [default to null]
**ParentAccountId** | **float64** | The ID of the parent account | [optional] [default to null]
**Name** | **string** | The name of the account | [default to null]
**Code** | **string** | The code of the account | [optional] [default to null]
**FiscalNumber** | **string** | The fiscal number of the account | [optional] [default to null]
**Address** | [***AccountAddressDto**](AccountAddressDto.md) |  | [optional] [default to null]
**PrimaryContactId** | **float64** | The user ID of the primary contact | [optional] [default to null]
**SecondaryContactId** | **float64** | The user ID of the secondary contact | [optional] [default to null]
**Archived** | **float64** | Whether the account is archived | [optional] [default to null]
**Limits** | [***AccountLimits**](AccountLimits.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

