# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryControllerUpdateVM**](VMsApi.md#InventoryControllerUpdateVM) | **Patch** /api/v2/vms/{vmId} | Updates VM information
[**VMMicroserviceControllerGetRemoteConsoleInfo**](VMsApi.md#VMMicroserviceControllerGetRemoteConsoleInfo) | **Get** /api/v2/vms/{vmId}/remote-console-info | Get Remote Console information
[**VMMicroserviceControllerGetVM**](VMsApi.md#VMMicroserviceControllerGetVM) | **Get** /api/v2/vms/{vmId} | Retrieves the VM information
[**VMMicroserviceControllerPowerStatusVMInstance**](VMsApi.md#VMMicroserviceControllerPowerStatusVMInstance) | **Get** /api/v2/vms/{vmId}/power-status | Retrieves the power status of the VM
[**VMMicroserviceControllerRebootVM**](VMsApi.md#VMMicroserviceControllerRebootVM) | **Post** /api/v2/vms/{vmId}/reboot | Reboots the VM
[**VMMicroserviceControllerShutdownVM**](VMsApi.md#VMMicroserviceControllerShutdownVM) | **Post** /api/v2/vms/{vmId}/shutdown | Shuts down the VM
[**VMMicroserviceControllerStartVM**](VMsApi.md#VMMicroserviceControllerStartVM) | **Post** /api/v2/vms/{vmId}/start | Starts the VM

# **InventoryControllerUpdateVM**
> Vm InventoryControllerUpdateVM(ctx, body, vmId)
Updates VM information

Updates VM information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVm**](UpdateVm.md)| The VM update object | 
  **vmId** | **float64**|  | 

### Return type

[**Vm**](VM.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VMMicroserviceControllerGetRemoteConsoleInfo**
> []RemoteConsoleInfoDto VMMicroserviceControllerGetRemoteConsoleInfo(ctx, vmId)
Get Remote Console information

Returns Remote Console information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **float64**|  | 

### Return type

[**[]RemoteConsoleInfoDto**](RemoteConsoleInfoDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VMMicroserviceControllerGetVM**
> VMMicroserviceControllerGetVM(ctx, vmId)
Retrieves the VM information

Retrieves the VM information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VMMicroserviceControllerPowerStatusVMInstance**
> VMMicroserviceControllerPowerStatusVMInstance(ctx, vmId)
Retrieves the power status of the VM

Retrieves the power status of the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VMMicroserviceControllerRebootVM**
> VMMicroserviceControllerRebootVM(ctx, vmId)
Reboots the VM

Reboots the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VMMicroserviceControllerShutdownVM**
> VMMicroserviceControllerShutdownVM(ctx, vmId)
Shuts down the VM

Shuts down the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VMMicroserviceControllerStartVM**
> VMMicroserviceControllerStartVM(ctx, vmId)
Starts the VM

Starts the VM

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

