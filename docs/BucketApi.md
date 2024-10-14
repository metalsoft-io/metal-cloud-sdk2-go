# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlueprintControllerCreateBucket**](BucketApi.md#BlueprintControllerCreateBucket) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets | Creates a Bucket
[**BlueprintControllerDeleteBucket**](BucketApi.md#BlueprintControllerDeleteBucket) | **Delete** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId} | Deletes a Bucket
[**BlueprintControllerGetBucket**](BucketApi.md#BlueprintControllerGetBucket) | **Get** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId} | Get Bucket information
[**BlueprintControllerGetBucketHosts**](BucketApi.md#BlueprintControllerGetBucketHosts) | **Get** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId}/hosts | Get the Hosts of Bucket
[**BlueprintControllerGetBuckets**](BucketApi.md#BlueprintControllerGetBuckets) | **Get** /api/v2/infrastructures/{infrastructureId}/buckets | Get all Buckets
[**BlueprintControllerGetIndividualBucket**](BucketApi.md#BlueprintControllerGetIndividualBucket) | **Get** /api/v2/buckets/{bucketId} | Get Bucket information
[**BlueprintControllerStartBucket**](BucketApi.md#BlueprintControllerStartBucket) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId}/actions/start | Starts a Bucket
[**BlueprintControllerStopBucket**](BucketApi.md#BlueprintControllerStopBucket) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId}/actions/stop | Stops a Bucket
[**BlueprintControllerUpdateBucket**](BucketApi.md#BlueprintControllerUpdateBucket) | **Patch** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId} | Updates Bucket information
[**BlueprintControllerUpdatebucketInstanceArrayHostsBulk**](BucketApi.md#BlueprintControllerUpdatebucketInstanceArrayHostsBulk) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId}/actions/modify-instance-array-hosts-bulk | Updates Instance Array Hosts on the Bucket

# **BlueprintControllerCreateBucket**
> Bucket BlueprintControllerCreateBucket(ctx, body, infrastructureId)
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

# **BlueprintControllerDeleteBucket**
> BlueprintControllerDeleteBucket(ctx, infrastructureId, bucketId)
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

# **BlueprintControllerGetBucket**
> Bucket BlueprintControllerGetBucket(ctx, infrastructureId, bucketId)
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

# **BlueprintControllerGetBucketHosts**
> BucketHosts BlueprintControllerGetBucketHosts(ctx, infrastructureId, bucketId)
Get the Hosts of Bucket

Returns the Hosts of Bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **bucketId** | **float64**|  | 

### Return type

[**BucketHosts**](BucketHosts.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerGetBuckets**
> []Bucket BlueprintControllerGetBuckets(ctx, infrastructureId)
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

# **BlueprintControllerGetIndividualBucket**
> BucketExtendedInfo BlueprintControllerGetIndividualBucket(ctx, bucketId)
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

# **BlueprintControllerStartBucket**
> BlueprintControllerStartBucket(ctx, infrastructureId, bucketId)
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

# **BlueprintControllerStopBucket**
> BlueprintControllerStopBucket(ctx, infrastructureId, bucketId)
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

# **BlueprintControllerUpdateBucket**
> Bucket BlueprintControllerUpdateBucket(ctx, body, infrastructureId, bucketId)
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

# **BlueprintControllerUpdatebucketInstanceArrayHostsBulk**
> BucketHosts BlueprintControllerUpdatebucketInstanceArrayHostsBulk(ctx, body, infrastructureId, bucketId)
Updates Instance Array Hosts on the Bucket

Updates Instance Array Hosts on the Bucket

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BucketHostsModifyBulk**](BucketHostsModifyBulk.md)| The Bucket Instance Array Hosts update object | 
  **infrastructureId** | **float64**|  | 
  **bucketId** | **float64**|  | 

### Return type

[**BucketHosts**](BucketHosts.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

