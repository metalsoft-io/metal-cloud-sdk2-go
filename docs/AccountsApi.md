# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAccount**](AccountsApi.md#ArchiveAccount) | **Post** /api/v2/accounts/{accountId}/actions/archive | Archive account
[**CreateAccount**](AccountsApi.md#CreateAccount) | **Post** /api/v2/accounts | Create account
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /api/v2/accounts/{accountId} | Get account by id
[**GetAccountLimits**](AccountsApi.md#GetAccountLimits) | **Get** /api/v2/accounts/{accountId}/limits | Get account limits
[**GetAccountUsers**](AccountsApi.md#GetAccountUsers) | **Get** /api/v2/accounts/{accountId}/users | Get users for account
[**GetAccounts**](AccountsApi.md#GetAccounts) | **Get** /api/v2/accounts | Get all accounts
[**UnarchiveAccount**](AccountsApi.md#UnarchiveAccount) | **Post** /api/v2/accounts/{accountId}/actions/unarchive | Unarchive account
[**UpdateAccount**](AccountsApi.md#UpdateAccount) | **Patch** /api/v2/accounts/{accountId} | Update account
[**UpdateAccountLimits**](AccountsApi.md#UpdateAccountLimits) | **Patch** /api/v2/accounts/{accountId}/limits | Update account limits

# **ArchiveAccount**
> AccountDto ArchiveAccount(ctx, accountId)
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

# **CreateAccount**
> AccountDto CreateAccount(ctx, body)
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

# **GetAccount**
> AccountDto GetAccount(ctx, accountId, optional)
Get account by id

Returns an account by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **float64**|  | 
 **optional** | ***AccountsApiGetAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountsApiGetAccountOpts struct
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

# **GetAccountLimits**
> AccountLimitsDto GetAccountLimits(ctx, accountId)
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

# **GetAccountUsers**
> []UserDto GetAccountUsers(ctx, accountId)
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

# **GetAccounts**
> []AccountDto GetAccounts(ctx, )
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

# **UnarchiveAccount**
> AccountDto UnarchiveAccount(ctx, accountId)
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

# **UpdateAccount**
> AccountDto UpdateAccount(ctx, body, accountId)
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

# **UpdateAccountLimits**
> AccountLimitsDto UpdateAccountLimits(ctx, body, accountId)
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

