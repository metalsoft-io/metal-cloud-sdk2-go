# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInfrastructureFileShare**](FileShareApi.md#CreateInfrastructureFileShare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares | Creates a File Share
[**DeleteFileShare**](FileShareApi.md#DeleteFileShare) | **Delete** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Deletes a File Share
[**GetFileShare**](FileShareApi.md#GetFileShare) | **Get** /api/v2/file-shares/{fileShareId} | Get File Share information
[**GetFileShareHosts**](FileShareApi.md#GetFileShareHosts) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/hosts | Get the Hosts of File Share
[**GetInfrastructureFileShare**](FileShareApi.md#GetInfrastructureFileShare) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Get File Share information
[**GetInfrastructureFileShares**](FileShareApi.md#GetInfrastructureFileShares) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares | Get all File Shares
[**StartFileShare**](FileShareApi.md#StartFileShare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/start | Starts a File Share
[**StopFileShare**](FileShareApi.md#StopFileShare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/stop | Stops a File Share
[**UpdateFileShare**](FileShareApi.md#UpdateFileShare) | **Patch** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Updates File Share information
[**UpdateFileShareInstanceArrayHostsBulk**](FileShareApi.md#UpdateFileShareInstanceArrayHostsBulk) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/modify-instance-array-hosts-bulk | Updates Instance Array Hosts on the File Share

# **CreateInfrastructureFileShare**
> FileShare CreateInfrastructureFileShare(ctx, body, infrastructureId)
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

# **DeleteFileShare**
> DeleteFileShare(ctx, infrastructureId, fileShareId)
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

# **GetFileShare**
> FileShareExtendedInfo GetFileShare(ctx, fileShareId)
Get File Share information

Returns File Share information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileShareId** | **float64**|  | 

### Return type

[**FileShareExtendedInfo**](FileShareExtendedInfo.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileShareHosts**
> FileShareHosts GetFileShareHosts(ctx, infrastructureId, fileShareId)
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

# **GetInfrastructureFileShare**
> FileShare GetInfrastructureFileShare(ctx, infrastructureId, fileShareId)
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

# **GetInfrastructureFileShares**
> []FileShare GetInfrastructureFileShares(ctx, infrastructureId)
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

# **StartFileShare**
> StartFileShare(ctx, infrastructureId, fileShareId)
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

# **StopFileShare**
> StopFileShare(ctx, infrastructureId, fileShareId)
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

# **UpdateFileShare**
> FileShare UpdateFileShare(ctx, body, infrastructureId, fileShareId)
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

# **UpdateFileShareInstanceArrayHostsBulk**
> FileShareHosts UpdateFileShareInstanceArrayHostsBulk(ctx, body, infrastructureId, fileShareId)
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

