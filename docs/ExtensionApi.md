# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveExtension**](ExtensionApi.md#ArchiveExtension) | **Post** /api/v2/extensions/{extensionId}/actions/archive | Archive published extension
[**CreateExtension**](ExtensionApi.md#CreateExtension) | **Post** /api/v2/extensions | Create extension
[**GetExtension**](ExtensionApi.md#GetExtension) | **Get** /api/v2/extensions/{extensionId} | Get details for an extension
[**GetExtensionDefinition**](ExtensionApi.md#GetExtensionDefinition) | **Get** /api/v2/extensions/{extensionId}/definition | Get definition for an extension
[**GetExtensions**](ExtensionApi.md#GetExtensions) | **Get** /api/v2/extensions | List available extensions
[**PublishExtension**](ExtensionApi.md#PublishExtension) | **Post** /api/v2/extensions/{extensionId}/actions/publish | Publish draft extension
[**UpdateExtension**](ExtensionApi.md#UpdateExtension) | **Patch** /api/v2/extensions/{extensionId} | Update extension

# **ArchiveExtension**
> ArchiveExtension(ctx, extensionId)
Archive published extension

Archives published extension.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateExtension**
> ExtensionDto CreateExtension(ctx, body)
Create extension

Returns details of the new extension

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExtensionDefinitionDto**](ExtensionDefinitionDto.md)| The extension details | 

### Return type

[**ExtensionDto**](ExtensionDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtension**
> ExtensionDto GetExtension(ctx, extensionId)
Get details for an extension

Returns details of the specified extension

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **float64**|  | 

### Return type

[**ExtensionDto**](ExtensionDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensionDefinition**
> ExtensionDefinitionDto GetExtensionDefinition(ctx, extensionId)
Get definition for an extension

Returns the definition of the specified extension

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **float64**|  | 

### Return type

[**ExtensionDefinitionDto**](ExtensionDefinitionDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExtensions**
> ExtensionListDto GetExtensions(ctx, optional)
List available extensions

Returns list of the available extensions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtensionApiGetExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtensionApiGetExtensionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter extension by status, name, label and description | 

### Return type

[**ExtensionListDto**](ExtensionListDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishExtension**
> PublishExtension(ctx, extensionId)
Publish draft extension

Publishes draft extension.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **extensionId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExtension**
> ExtensionDto UpdateExtension(ctx, body, extensionId)
Update extension

Returns details of the updated extension

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExtensionDefinitionDto**](ExtensionDefinitionDto.md)| The extension details | 
  **extensionId** | **float64**|  | 

### Return type

[**ExtensionDto**](ExtensionDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

