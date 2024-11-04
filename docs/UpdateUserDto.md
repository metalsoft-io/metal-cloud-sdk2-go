# UpdateUserDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The new display name of the user | [optional] [default to null]
**Password** | [***UserUpdatePasswordDto**](UserUpdatePasswordDto.md) |  | [optional] [default to null]
**Email** | **string** | The new login e-mail of the user | [optional] [default to null]
**RedirectUrl** | **string** | The redirect URL for the user | [optional] [default to null]
**AESKey** | **string** | The AES key for the user | [optional] [default to null]
**EmailStatus** | **string** | The new e-mail status of the user | [optional] [default to null]
**Language** | **string** | The new language of the user | [optional] [default to null]
**Brand** | **string** | The new brand of the user | [optional] [default to null]
**IsBrandManager** | **bool** | The new brand manager status of the user | [optional] [default to null]
**Delegates** | [**[]UserDelegateDto**](UserDelegateDto.md) | The new user delegates of the user. | [optional] [default to null]
**IsBlocked** | **bool** | The new blocked status of the user | [optional] [default to null]
**Suspension** | [***UserSuspendDto**](UserSuspendDto.md) |  | [optional] [default to null]
**AccessLevel** | **string** | The new access level of the user | [optional] [default to null]
**Promotions** | [**[]UserPromotionDto**](UserPromotionDto.md) | The new promotions of the user. | [optional] [default to null]
**ExperimentalTags** | [**[]UserExperimentalTagDto**](UserExperimentalTagDto.md) | The new experimental tags of the user. | [optional] [default to null]
**CustomPrices** | [***interface{}**](interface{}.md) | The new custom prices of the user | [optional] [default to null]
**Permissions** | [***UserResourcePermissionsDto**](UserResourcePermissionsDto.md) |  | [optional] [default to null]
**IsTestAccount** | **bool** | Whether the user account is a test one. | [optional] [default to null]
**IsBillable** | **bool** | Whether the user account is billable. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

