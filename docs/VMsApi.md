# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVM**](VMsApi.md#GetVM) | **Get** /api/v2/vms/{vmId} | Retrieves the VM information
[**GetVMPowerStatus**](VMsApi.md#GetVMPowerStatus) | **Get** /api/v2/vms/{vmId}/power-status | Retrieves the power status of the VM
[**GetVMRemoteConsoleInfo**](VMsApi.md#GetVMRemoteConsoleInfo) | **Get** /api/v2/vms/{vmId}/remote-console-info | Get Remote Console information
[**RebootVM**](VMsApi.md#RebootVM) | **Post** /api/v2/vms/{vmId}/reboot | Reboots the VM
[**ShutdownVM**](VMsApi.md#ShutdownVM) | **Post** /api/v2/vms/{vmId}/shutdown | Shuts down the VM
[**StartVM**](VMsApi.md#StartVM) | **Post** /api/v2/vms/{vmId}/start | Starts the VM
[**UpdateVM**](VMsApi.md#UpdateVM) | **Patch** /api/v2/vms/{vmId} | Updates VM information

# **GetVM**
> GetVM(ctx, vmId)
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

# **GetVMPowerStatus**
> GetVMPowerStatus(ctx, vmId)
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

# **GetVMRemoteConsoleInfo**
> []RemoteConsoleInfoDto GetVMRemoteConsoleInfo(ctx, vmId)
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

# **RebootVM**
> RebootVM(ctx, vmId)
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

# **ShutdownVM**
> ShutdownVM(ctx, vmId)
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

# **StartVM**
> StartVM(ctx, vmId)
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

# **UpdateVM**
> Vm UpdateVM(ctx, body, vmId)
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

