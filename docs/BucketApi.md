# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInfrastructureBucket**](BucketApi.md#CreateInfrastructureBucket) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets | Creates a Bucket
[**DeleteBucket**](BucketApi.md#DeleteBucket) | **Delete** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId} | Deletes a Bucket
[**GetBucket**](BucketApi.md#GetBucket) | **Get** /api/v2/buckets/{bucketId} | Get Bucket information
[**GetInfrastructureBucket**](BucketApi.md#GetInfrastructureBucket) | **Get** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId} | Get Bucket information
[**GetInfrastructureBuckets**](BucketApi.md#GetInfrastructureBuckets) | **Get** /api/v2/infrastructures/{infrastructureId}/buckets | Get all Buckets
[**StartBucket**](BucketApi.md#StartBucket) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId}/actions/start | Starts a Bucket
[**StopBucket**](BucketApi.md#StopBucket) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId}/actions/stop | Stops a Bucket
[**UpdateBucket**](BucketApi.md#UpdateBucket) | **Patch** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId} | Updates Bucket information

# **CreateInfrastructureBucket**
> Bucket CreateInfrastructureBucket(ctx, body, infrastructureId)
Creates a Bucket

Creates a Bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBucket**](CreateBucket.md)| The Bucket create object | 
  **infrastructureId** | **float64**|  | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBucket**
> DeleteBucket(ctx, infrastructureId, bucketId)
Deletes a Bucket

Deletes a Bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **bucketId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBucket**
> BucketExtendedInfo GetBucket(ctx, bucketId)
Get Bucket information

Returns Bucket information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bucketId** | **float64**|  | 

### Return type

[**BucketExtendedInfo**](BucketExtendedInfo.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfrastructureBucket**
> Bucket GetInfrastructureBucket(ctx, infrastructureId, bucketId)
Get Bucket information

Returns Bucket information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **bucketId** | **float64**|  | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfrastructureBuckets**
> []Bucket GetInfrastructureBuckets(ctx, infrastructureId)
Get all Buckets

Returns list of all Buckets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 

### Return type

[**[]Bucket**](Bucket.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartBucket**
> StartBucket(ctx, infrastructureId, bucketId)
Starts a Bucket

Starts a Bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **bucketId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopBucket**
> StopBucket(ctx, infrastructureId, bucketId)
Stops a Bucket

Stops a Bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **bucketId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBucket**
> Bucket UpdateBucket(ctx, body, infrastructureId, bucketId)
Updates Bucket information

Updates Bucket information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBucket**](UpdateBucket.md)| The Bucket update object | 
  **infrastructureId** | **float64**|  | 
  **bucketId** | **float64**|  | 

### Return type

[**Bucket**](Bucket.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

