# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtensionsControllerArchiveExtension**](ExtensionApi.md#ExtensionsControllerArchiveExtension) | **Post** /api/v2/extensions/{extensionId}/actions/archive | Archive published extension
[**ExtensionsControllerCreateExtension**](ExtensionApi.md#ExtensionsControllerCreateExtension) | **Post** /api/v2/extensions | Create extension
[**ExtensionsControllerGetExtension**](ExtensionApi.md#ExtensionsControllerGetExtension) | **Get** /api/v2/extensions/{extensionId} | Get details for an extension
[**ExtensionsControllerListExtensions**](ExtensionApi.md#ExtensionsControllerListExtensions) | **Get** /api/v2/extensions | List available extensions
[**ExtensionsControllerPublishExtension**](ExtensionApi.md#ExtensionsControllerPublishExtension) | **Post** /api/v2/extensions/{extensionId}/actions/publish | Publish draft extension
[**ExtensionsControllerUpdateExtension**](ExtensionApi.md#ExtensionsControllerUpdateExtension) | **Patch** /api/v2/extensions/{extensionId} | Update extension

# **ExtensionsControllerArchiveExtension**
> ExtensionsControllerArchiveExtension(ctx, extensionId)
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

# **ExtensionsControllerCreateExtension**
> ExtensionDto ExtensionsControllerCreateExtension(ctx, body)
Create extension

Returns details of the new extension

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateExtensionDto**](CreateExtensionDto.md)| The extension details | 

### Return type

[**ExtensionDto**](ExtensionDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtensionsControllerGetExtension**
> ExtensionDto ExtensionsControllerGetExtension(ctx, extensionId)
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

# **ExtensionsControllerListExtensions**
> ExtensionListDto ExtensionsControllerListExtensions(ctx, optional)
List available extensions

Returns list of the available extensions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtensionApiExtensionsControllerListExtensionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtensionApiExtensionsControllerListExtensionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter by extension status | 

### Return type

[**ExtensionListDto**](ExtensionListDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtensionsControllerPublishExtension**
> ExtensionsControllerPublishExtension(ctx, extensionId)
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

# **ExtensionsControllerUpdateExtension**
> ExtensionDto ExtensionsControllerUpdateExtension(ctx, body, extensionId)
Update extension

Returns details of the updated extension

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateExtensionDto**](UpdateExtensionDto.md)| The extension details | 
  **extensionId** | **float64**|  | 

### Return type

[**ExtensionDto**](ExtensionDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

