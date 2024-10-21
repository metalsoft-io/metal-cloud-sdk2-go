# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExtensionInstance**](ExtensionInstanceApi.md#CreateExtensionInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/extension-instances | Add extension instance to an infrastructure
[**DeleteExtensionInstance**](ExtensionInstanceApi.md#DeleteExtensionInstance) | **Delete** /api/v2/extension-instances/{extensionInstanceId} | Delete extension instance
[**GetExtensionInstance**](ExtensionInstanceApi.md#GetExtensionInstance) | **Get** /api/v2/extension-instances/{extensionInstanceId} | Get extension instance details
[**GetExtensionInstances**](ExtensionInstanceApi.md#GetExtensionInstances) | **Get** /api/v2/extension-instances | Get extension instances list
[**UpdateExtensionInstance**](ExtensionInstanceApi.md#UpdateExtensionInstance) | **Patch** /api/v2/extension-instances/{extensionInstanceId} | Update extension instance configuration

# **CreateExtensionInstance**
> ExtensionInstanceDto CreateExtensionInstance(ctx, body, infrastructureId)
Add extension instance to an infrastructure

Adds extension instance to the specified infrastructure. Note that a infrastructure deploy is needed for the changes to take effect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateExtensionInstanceDto**](CreateExtensionInstanceDto.md)| The extension instance to create | 
  **infrastructureId** | **float64**|  | 

### Return type

[**ExtensionInstanceDto**](ExtensionInstanceDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExtensionInstance**
> DeleteExtensionInstance(ctx, extensionInstanceId)
Delete extension instance

Deletes the specified extension instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionInstanceId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionInstance**
> ExtensionInstanceDto GetExtensionInstance(ctx, extensionInstanceId)
Get extension instance details

Returns the details of the specified extension instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionInstanceId** | **float64**|  | 

### Return type

[**ExtensionInstanceDto**](ExtensionInstanceDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionInstances**
> ExtensionInstanceListDto GetExtensionInstances(ctx, optional)
Get extension instances list

Returns list of extension instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtensionInstanceApiGetExtensionInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtensionInstanceApiGetExtensionInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**| Page number used in pagination, starts with &#x60;0&#x60; | 
 **limit** | **optional.Float64**| Number of items per page, default value &#x60;20&#x60; | 
 **sortBy** | **optional.String**| Sort by field and direction, example: &#x60;id:ASC&#x60; | 
 **filter** | [**optional.Interface of interface{}**](.md)| Filter expression for a field, see [nestjs paginate documentation](https://github.com/ppetzold/nestjs-paginate#filters).&lt;br /&gt; &lt;br /&gt; Use &#x60;\&quot;filter.*\&quot;&#x60; as a way to specify the field to filter on, see the example bellow: &lt;br /&gt; &lt;br /&gt;  | 

### Return type

[**ExtensionInstanceListDto**](ExtensionInstanceListDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExtensionInstance**
> ExtensionInstanceDto UpdateExtensionInstance(ctx, body, extensionInstanceId)
Update extension instance configuration

Updates the specified extension instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateExtensionInstanceDto**](UpdateExtensionInstanceDto.md)| The extension instance changes | 
  **extensionInstanceId** | **float64**|  | 

### Return type

[**ExtensionInstanceDto**](ExtensionInstanceDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

