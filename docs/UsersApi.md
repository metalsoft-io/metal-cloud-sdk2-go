# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersControllerArchiveUser**](UsersApi.md#UsersControllerArchiveUser) | **Post** /api/v2/users/{userId}/actions/archive | Archive user
[**UsersControllerChangeUserAccount**](UsersApi.md#UsersControllerChangeUserAccount) | **Post** /api/v2/users/{userId}/actions/change-account | Change account for user
[**UsersControllerCreateUser**](UsersApi.md#UsersControllerCreateUser) | **Post** /api/v2/users | Creates a user
[**UsersControllerGetUser**](UsersApi.md#UsersControllerGetUser) | **Get** /api/v2/users/{userId} | Get user
[**UsersControllerGetUserLimits**](UsersApi.md#UsersControllerGetUserLimits) | **Get** /api/v2/users/{userId}/limits | Get user limits
[**UsersControllerGetUsers**](UsersApi.md#UsersControllerGetUsers) | **Get** /api/v2/users | Get users
[**UsersControllerUnarchiveUser**](UsersApi.md#UsersControllerUnarchiveUser) | **Post** /api/v2/users/{userId}/actions/unarchive | Unarchive user
[**UsersControllerUpdateUser**](UsersApi.md#UsersControllerUpdateUser) | **Patch** /api/v2/users/{userId} | Update user
[**UsersControllerUpdateUserLimits**](UsersApi.md#UsersControllerUpdateUserLimits) | **Patch** /api/v2/users/{userId}/limits | Update user limits

# **UsersControllerArchiveUser**
> UserDto UsersControllerArchiveUser(ctx, userId)
Archive user

Archives a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerChangeUserAccount**
> UserDto UsersControllerChangeUserAccount(ctx, body, userId)
Change account for user

Changes account for user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeUserAccountDto**](ChangeUserAccountDto.md)| The new account id | 
  **userId** | **float64**|  | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerCreateUser**
> UserDto UsersControllerCreateUser(ctx, body)
Creates a user

Creates a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateUserDto**](CreateUserDto.md)| The user to create | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerGetUser**
> UserDto UsersControllerGetUser(ctx, userId, optional)
Get user

Returns a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 
 **optional** | ***UsersApiUsersControllerGetUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersControllerGetUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursion** | **optional.Float64**| The recursion level of the displayed details. Default is 0. | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerGetUserLimits**
> UserLimitsDto UsersControllerGetUserLimits(ctx, userId)
Get user limits

Returns the limits of a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**UserLimitsDto**](UserLimitsDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerGetUsers**
> []UserDto UsersControllerGetUsers(ctx, )
Get users

Returns a list of users

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserDto**](UserDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerUnarchiveUser**
> UserDto UsersControllerUnarchiveUser(ctx, userId)
Unarchive user

Unarchives a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerUpdateUser**
> UserDto UsersControllerUpdateUser(ctx, body, userId)
Update user

Updates a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserDto**](UpdateUserDto.md)| The user updates | 
  **userId** | **float64**|  | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersControllerUpdateUserLimits**
> UserLimitsDto UsersControllerUpdateUserLimits(ctx, body, userId)
Update user limits

Updates the limits of a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserLimitsDto**](UserLimitsDto.md)| The new limits | 
  **userId** | **float64**|  | 

### Return type

[**UserLimitsDto**](UserLimitsDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

