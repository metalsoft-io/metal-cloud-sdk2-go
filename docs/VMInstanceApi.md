# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlueprintControllerApplyVMTypeOnVMInstance**](VMInstanceApi.md#BlueprintControllerApplyVMTypeOnVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance
[**BlueprintControllerCreateVMInstance**](VMInstanceApi.md#BlueprintControllerCreateVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances | Creates a VM Instance
[**BlueprintControllerDeleteVMInstance**](VMInstanceApi.md#BlueprintControllerDeleteVMInstance) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Deletes a VM Instance
[**BlueprintControllerGetVMInstance**](VMInstanceApi.md#BlueprintControllerGetVMInstance) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Get VM Instance information
[**BlueprintControllerPowerStatusVMInstance**](VMInstanceApi.md#BlueprintControllerPowerStatusVMInstance) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/power-status | Retrieves the power status of the VM Instance
[**BlueprintControllerRebootVMInstance**](VMInstanceApi.md#BlueprintControllerRebootVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/reboot | Reboots the VM Instance
[**BlueprintControllerShutdownVMInstance**](VMInstanceApi.md#BlueprintControllerShutdownVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/shutdown | Shuts down the VM Instance
[**BlueprintControllerStartVMInstance**](VMInstanceApi.md#BlueprintControllerStartVMInstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/start | Starts the VM Instance
[**BlueprintControllerUpdateVMInstance**](VMInstanceApi.md#BlueprintControllerUpdateVMInstance) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Updates VM Instance information

# **BlueprintControllerApplyVMTypeOnVMInstance**
> VmInstanceDto BlueprintControllerApplyVMTypeOnVMInstance(ctx, infrastructureId, vmInstanceId, vmTypeId)
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

[**VmInstanceDto**](VMInstanceDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerCreateVMInstance**
> VmInstanceDto BlueprintControllerCreateVMInstance(ctx, body, infrastructureId)
Creates a VM Instance

Creates a VM Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVmInstanceDto**](CreateVmInstanceDto.md)| The VM Instance create object | 
  **infrastructureId** | **float64**|  | 

### Return type

[**VmInstanceDto**](VMInstanceDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerDeleteVMInstance**
> BlueprintControllerDeleteVMInstance(ctx, infrastructureId, vmInstanceId)
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

# **BlueprintControllerGetVMInstance**
> VmInstanceDto BlueprintControllerGetVMInstance(ctx, infrastructureId, vmInstanceId)
Get VM Instance information

Returns VM Instance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 

### Return type

[**VmInstanceDto**](VMInstanceDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlueprintControllerPowerStatusVMInstance**
> BlueprintControllerPowerStatusVMInstance(ctx, infrastructureId, vmInstanceId)
Retrieves the power status of the VM Instance

Retrieves the power status of the VM Instance

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

# **BlueprintControllerRebootVMInstance**
> BlueprintControllerRebootVMInstance(ctx, infrastructureId, vmInstanceId)
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

# **BlueprintControllerShutdownVMInstance**
> BlueprintControllerShutdownVMInstance(ctx, infrastructureId, vmInstanceId)
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

# **BlueprintControllerStartVMInstance**
> BlueprintControllerStartVMInstance(ctx, infrastructureId, vmInstanceId)
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

# **BlueprintControllerUpdateVMInstance**
> VmInstanceDto BlueprintControllerUpdateVMInstance(ctx, body, infrastructureId, vmInstanceId)
Updates VM Instance information

Updates VM Instance information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmInstanceDto**](UpdateVmInstanceDto.md)| The VM Instance update object | 
  **infrastructureId** | **float64**|  | 
  **vmInstanceId** | **float64**|  | 

### Return type

[**VmInstanceDto**](VMInstanceDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

