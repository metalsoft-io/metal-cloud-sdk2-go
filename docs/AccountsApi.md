# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsControllerArchiveAccount**](AccountsApi.md#AccountsControllerArchiveAccount) | **Post** /api/v2/accounts/{accountId}/actions/archive | Archive account
[**AccountsControllerCreateAccount**](AccountsApi.md#AccountsControllerCreateAccount) | **Post** /api/v2/accounts | Create account
[**AccountsControllerGetAccount**](AccountsApi.md#AccountsControllerGetAccount) | **Get** /api/v2/accounts/{accountId} | Get account by id
[**AccountsControllerGetAccountLimits**](AccountsApi.md#AccountsControllerGetAccountLimits) | **Get** /api/v2/accounts/{accountId}/limits | Get account limits
[**AccountsControllerGetAccounts**](AccountsApi.md#AccountsControllerGetAccounts) | **Get** /api/v2/accounts | Get all accounts
[**AccountsControllerGetUsersForAccount**](AccountsApi.md#AccountsControllerGetUsersForAccount) | **Get** /api/v2/accounts/{accountId}/users | Get users for account
[**AccountsControllerUnarchiveAccount**](AccountsApi.md#AccountsControllerUnarchiveAccount) | **Post** /api/v2/accounts/{accountId}/actions/unarchive | Unarchive account
[**AccountsControllerUpdateAccount**](AccountsApi.md#AccountsControllerUpdateAccount) | **Patch** /api/v2/accounts/{accountId} | Update account
[**AccountsControllerUpdateAccountLimits**](AccountsApi.md#AccountsControllerUpdateAccountLimits) | **Patch** /api/v2/accounts/{accountId}/limits | Update account limits

# **AccountsControllerArchiveAccount**
> AccountDto AccountsControllerArchiveAccount(ctx, accountId)
Archive account

Archives an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **float64**|  | 

### Return type

[**AccountDto**](AccountDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsControllerCreateAccount**
> AccountDto AccountsControllerCreateAccount(ctx, body)
Create account

Creates an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAccountDto**](CreateAccountDto.md)| The account to create | 

### Return type

[**AccountDto**](AccountDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsControllerGetAccount**
> AccountDto AccountsControllerGetAccount(ctx, accountId, optional)
Get account by id

Returns an account by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **float64**|  | 
 **optional** | ***AccountsApiAccountsControllerGetAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiAccountsControllerGetAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursion** | **optional.Float64**| The recursion level of the displayed details. Default is 0. | 

### Return type

[**AccountDto**](AccountDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsControllerGetAccountLimits**
> AccountLimitsDto AccountsControllerGetAccountLimits(ctx, accountId)
Get account limits

Returns the limits of an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **float64**|  | 

### Return type

[**AccountLimitsDto**](AccountLimitsDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsControllerGetAccounts**
> []AccountDto AccountsControllerGetAccounts(ctx, )
Get all accounts

Returns all accounts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AccountDto**](AccountDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsControllerGetUsersForAccount**
> []UserDto AccountsControllerGetUsersForAccount(ctx, accountId)
Get users for account

Returns users for an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **float64**|  | 

### Return type

[**[]UserDto**](UserDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsControllerUnarchiveAccount**
> AccountDto AccountsControllerUnarchiveAccount(ctx, accountId)
Unarchive account

Unarchives an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **float64**|  | 

### Return type

[**AccountDto**](AccountDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsControllerUpdateAccount**
> AccountDto AccountsControllerUpdateAccount(ctx, body, accountId)
Update account

Updates an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAccountDto**](UpdateAccountDto.md)| The account updates | 
  **accountId** | **float64**|  | 

### Return type

[**AccountDto**](AccountDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AccountsControllerUpdateAccountLimits**
> AccountLimitsDto AccountsControllerUpdateAccountLimits(ctx, body, accountId)
Update account limits

Updates the limits of an account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccountLimitsDto**](AccountLimitsDto.md)| The account limits updates | 
  **accountId** | **float64**|  | 

### Return type

[**AccountLimitsDto**](AccountLimitsDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

