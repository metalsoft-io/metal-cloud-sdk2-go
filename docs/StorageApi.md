# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlueprintControllerGetStorageBuckets**](StorageApi.md#BlueprintControllerGetStorageBuckets) | **Get** /api/v2/storages/{storageId}/buckets | Get all Buckets linked to the specified storage
[**BlueprintControllerGetStorageFileShares**](StorageApi.md#BlueprintControllerGetStorageFileShares) | **Get** /api/v2/storages/{storageId}/file-shares | Get all File Shares linked to the specified storage
[**InventoryControllerCreateStorage**](StorageApi.md#InventoryControllerCreateStorage) | **Post** /api/v2/storages | Creates a Storage
[**InventoryControllerGetStorage**](StorageApi.md#InventoryControllerGetStorage) | **Get** /api/v2/storages/{storageId} | Retrieves a Storage
[**InventoryControllerPatchStorage**](StorageApi.md#InventoryControllerPatchStorage) | **Patch** /api/v2/storages/{storageId} | Updates a Storage

# **BlueprintControllerGetStorageBuckets**
> BucketList BlueprintControllerGetStorageBuckets(ctx, storageId)
Get all Buckets linked to the specified storage

Returns list of all Buckets linked to the specified Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storageId** | **float64**|  | 

### Return type

[**BucketList**](BucketList.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerGetStorageFileShares**
> FileShareList BlueprintControllerGetStorageFileShares(ctx, storageId)
Get all File Shares linked to the specified storage

Returns list of all File Shares linked to the specified Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storageId** | **float64**|  | 

### Return type

[**FileShareList**](FileShareList.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerCreateStorage**
> StorageRegistrationResponse InventoryControllerCreateStorage(ctx, body)
Creates a Storage

Creates a Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateStorage**](CreateStorage.md)| The Storage create object | 

### Return type

[**StorageRegistrationResponse**](StorageRegistrationResponse.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerGetStorage**
> StorageRegistrationResponse InventoryControllerGetStorage(ctx, storageId)
Retrieves a Storage

Retrieves a Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **storageId** | **float64**|  | 

### Return type

[**StorageRegistrationResponse**](StorageRegistrationResponse.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryControllerPatchStorage**
> InventoryControllerPatchStorage(ctx, body, storageId)
Updates a Storage

Updates a Storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatchStorage**](PatchStorage.md)| The Storage update object | 
  **storageId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

