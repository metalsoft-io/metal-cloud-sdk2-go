# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlueprintControllerCreateFileShare**](FileShareApi.md#BlueprintControllerCreateFileShare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares | Creates a File Share
[**BlueprintControllerDeleteFileShare**](FileShareApi.md#BlueprintControllerDeleteFileShare) | **Delete** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Deletes a File Share
[**BlueprintControllerGetFileShare**](FileShareApi.md#BlueprintControllerGetFileShare) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Get File Share information
[**BlueprintControllerGetFileShareHosts**](FileShareApi.md#BlueprintControllerGetFileShareHosts) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/hosts | Get the Hosts of File Share
[**BlueprintControllerGetFileShares**](FileShareApi.md#BlueprintControllerGetFileShares) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares | Get all File Shares
[**BlueprintControllerStartFileShare**](FileShareApi.md#BlueprintControllerStartFileShare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/start | Starts a File Share
[**BlueprintControllerStopFileShare**](FileShareApi.md#BlueprintControllerStopFileShare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/stop | Stops a File Share
[**BlueprintControllerUpdateFileShare**](FileShareApi.md#BlueprintControllerUpdateFileShare) | **Patch** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Updates File Share information
[**BlueprintControllerUpdateFileShareInstanceArrayHostsBulk**](FileShareApi.md#BlueprintControllerUpdateFileShareInstanceArrayHostsBulk) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/modify-instance-array-hosts-bulk | Updates Instance Array Hosts on the File Share

# **BlueprintControllerCreateFileShare**
> FileShare BlueprintControllerCreateFileShare(ctx, body, infrastructureId)
Creates a File Share

Creates a File Share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateFileShare**](CreateFileShare.md)| The File Share create object | 
  **infrastructureId** | **float64**|  | 

### Return type

[**FileShare**](FileShare.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerDeleteFileShare**
> BlueprintControllerDeleteFileShare(ctx, infrastructureId, fileShareId)
Deletes a File Share

Deletes a File Share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **fileShareId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerGetFileShare**
> FileShare BlueprintControllerGetFileShare(ctx, infrastructureId, fileShareId)
Get File Share information

Returns File Share information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **fileShareId** | **float64**|  | 

### Return type

[**FileShare**](FileShare.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerGetFileShareHosts**
> FileShareHosts BlueprintControllerGetFileShareHosts(ctx, infrastructureId, fileShareId)
Get the Hosts of File Share

Returns the Hosts of File Share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **fileShareId** | **float64**|  | 

### Return type

[**FileShareHosts**](FileShareHosts.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerGetFileShares**
> []FileShare BlueprintControllerGetFileShares(ctx, infrastructureId)
Get all File Shares

Returns list of all File Shares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 

### Return type

[**[]FileShare**](FileShare.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerStartFileShare**
> BlueprintControllerStartFileShare(ctx, infrastructureId, fileShareId)
Starts a File Share

Starts a File Share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **fileShareId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerStopFileShare**
> BlueprintControllerStopFileShare(ctx, infrastructureId, fileShareId)
Stops a File Share

Stops a File Share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **fileShareId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerUpdateFileShare**
> FileShare BlueprintControllerUpdateFileShare(ctx, body, infrastructureId, fileShareId)
Updates File Share information

Updates File Share information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateFileShare**](UpdateFileShare.md)| The File Share update object | 
  **infrastructureId** | **float64**|  | 
  **fileShareId** | **float64**|  | 

### Return type

[**FileShare**](FileShare.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerUpdateFileShareInstanceArrayHostsBulk**
> FileShareHosts BlueprintControllerUpdateFileShareInstanceArrayHostsBulk(ctx, body, infrastructureId, fileShareId)
Updates Instance Array Hosts on the File Share

Updates Instance Array Hosts on the File Share

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FileShareHostsModifyBulk**](FileShareHostsModifyBulk.md)| The File Share Instance Array Hosts update object | 
  **infrastructureId** | **float64**|  | 
  **fileShareId** | **float64**|  | 

### Return type

[**FileShareHosts**](FileShareHosts.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

