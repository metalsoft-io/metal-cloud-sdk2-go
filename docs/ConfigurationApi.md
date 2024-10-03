# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigController1GetConfiguration**](ConfigurationApi.md#ConfigController1GetConfiguration) | **Get** /api/v2/config | Get configuration

# **ConfigController1GetConfiguration**
> interface{} ConfigController1GetConfiguration(ctx, optional)
Get configuration

Returns a configuration object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigurationApiConfigController1GetConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigurationApiConfigController1GetConfigurationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter to be applied. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

