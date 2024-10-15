# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListProviders**](SecurityApi.md#ListProviders) | **Get** /api/v2/authentication/providers | Get available authentication providers
[**UpdateProvider**](SecurityApi.md#UpdateProvider) | **Patch** /api/v2/authentication/providers/{name} | Updates authentication provider

# **ListProviders**
> []AuthenticationProvider ListProviders(ctx, )
Get available authentication providers

Returns the available authentication providers

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]AuthenticationProvider**](AuthenticationProvider.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProvider**
> []AuthenticationProvider UpdateProvider(ctx, body, name)
Updates authentication provider

Updates authentication provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthenticationProviderUpdate**](AuthenticationProviderUpdate.md)| The provider updates | 
  **name** | **string**|  | 

### Return type

[**[]AuthenticationProvider**](AuthenticationProvider.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

