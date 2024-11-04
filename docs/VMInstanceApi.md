# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyVMTypeOnVMInstance**](VMInstanceApi.md#ApplyVMTypeOnVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance
[**CreateVMInstance**](VMInstanceApi.md#CreateVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances | Creates a VM Instance
[**DeleteVMInstance**](VMInstanceApi.md#DeleteVMInstance) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Deletes a VM Instance
[**GetVMInstance**](VMInstanceApi.md#GetVMInstance) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Get VM Instance information
[**GetVMInstancePowerStatus**](VMInstanceApi.md#GetVMInstancePowerStatus) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/power-status | Retrieves the power status of the VM Instance
[**RebootVMInstance**](VMInstanceApi.md#RebootVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/reboot | Reboots the VM Instance
[**ShutdownVMInstance**](VMInstanceApi.md#ShutdownVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/shutdown | Shuts down the VM Instance
[**StartVMInstance**](VMInstanceApi.md#StartVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/start | Starts the VM Instance
[**UpdateVMInstance**](VMInstanceApi.md#UpdateVMInstance) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Updates VM Instance information

# **ApplyVMTypeOnVMInstance**
> VmInstance ApplyVMTypeOnVMInstance(ctx, infrastructureId, vmInstanceId, vmTypeId)
Applies a VM Type to a VM Instance

Applies a VM Type to a VM Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 
  **vmTypeId** | **float64**|  | 

### Return type

[**VmInstance**](VMInstance.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVMInstance**
> VmInstance CreateVMInstance(ctx, body, infrastructureId)
Creates a VM Instance

Creates a VM Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVmInstance**](CreateVmInstance.md)| The VM Instance create object | 
  **infrastructureId** | **float64**|  | 

### Return type

[**VmInstance**](VMInstance.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVMInstance**
> DeleteVMInstance(ctx, infrastructureId, vmInstanceId)
Deletes a VM Instance

Deletes a VM Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMInstance**
> VmInstance GetVMInstance(ctx, infrastructureId, vmInstanceId)
Get VM Instance information

Returns VM Instance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 

### Return type

[**VmInstance**](VMInstance.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMInstancePowerStatus**
> string GetVMInstancePowerStatus(ctx, infrastructureId, vmInstanceId)
Retrieves the power status of the VM Instance

Retrieves the power status of the VM Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 

### Return type

**string**

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RebootVMInstance**
> RebootVMInstance(ctx, infrastructureId, vmInstanceId)
Reboots the VM Instance

Reboots the VM Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShutdownVMInstance**
> ShutdownVMInstance(ctx, infrastructureId, vmInstanceId)
Shuts down the VM Instance

Shuts down the VM Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartVMInstance**
> StartVMInstance(ctx, infrastructureId, vmInstanceId)
Starts the VM Instance

Starts the VM Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVMInstance**
> VmInstance UpdateVMInstance(ctx, body, infrastructureId, vmInstanceId)
Updates VM Instance information

Updates VM Instance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmInstance**](UpdateVmInstance.md)| The VM Instance update object | 
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 

### Return type

[**VmInstance**](VMInstance.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

