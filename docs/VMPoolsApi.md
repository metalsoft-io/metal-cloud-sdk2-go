# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVMPool**](VMPoolsApi.md#CreateVMPool) | **Post** /api/v2/vm-pools | Creates a VM Pool
[**DeleteVMPool**](VMPoolsApi.md#DeleteVMPool) | **Delete** /api/v2/vm-pools/{vmPoolId} | Deletes a VM Pool
[**GetVMPool**](VMPoolsApi.md#GetVMPool) | **Get** /api/v2/vm-pools/{vmPoolId} | Get VM Pool information
[**GetVMPoolClusterHost**](VMPoolsApi.md#GetVMPoolClusterHost) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId} | Retrieves a VM Cluster Host
[**GetVMPoolClusterHostInterface**](VMPoolsApi.md#GetVMPoolClusterHostInterface) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Retrieves a VM Cluster Host Interface
[**GetVMPoolClusterHostInterfaces**](VMPoolsApi.md#GetVMPoolClusterHostInterfaces) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces | Retrieves a list of VM Cluster Host Interfaces
[**GetVMPoolClusterHostVMs**](VMPoolsApi.md#GetVMPoolClusterHostVMs) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/vms | Retrieves a list of VM Cluster Host VMs
[**GetVMPoolClusterHosts**](VMPoolsApi.md#GetVMPoolClusterHosts) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts | Get list of VM Cluster Hosts linked to the VM Pool
[**GetVMPoolVMs**](VMPoolsApi.md#GetVMPoolVMs) | **Get** /api/v2/vm-pools/{vmPoolId}/vms | Returns all VMs linked to the VM Pool
[**GetVMPools**](VMPoolsApi.md#GetVMPools) | **Get** /api/v2/vm-pools | Get all VM Pools
[**UpdateVMPool**](VMPoolsApi.md#UpdateVMPool) | **Patch** /api/v2/vm-pools/{vmPoolId} | Updates VM Pool information
[**UpdateVMPoolClusterHostInterface**](VMPoolsApi.md#UpdateVMPoolClusterHostInterface) | **Patch** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Updates a VM Cluster Host Interface

# **CreateVMPool**
> VmPool CreateVMPool(ctx, body)
Creates a VM Pool

Creates a VM Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVmPool**](CreateVmPool.md)| The VM Pool create object | 

### Return type

[**VmPool**](VMPool.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVMPool**
> DeleteVMPool(ctx, vmPoolId)
Deletes a VM Pool

Deletes a VM Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMPool**
> VmPool GetVMPool(ctx, vmPoolId)
Get VM Pool information

Returns VM Pool information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 

### Return type

[**VmPool**](VMPool.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMPoolClusterHost**
> VmPoolHosts GetVMPoolClusterHost(ctx, vmPoolId, vmPoolClusterHostId)
Retrieves a VM Cluster Host

Returns a VM Cluster Host

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 

### Return type

[**VmPoolHosts**](VMPoolHosts.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMPoolClusterHostInterface**
> VmPoolHostInterfaces GetVMPoolClusterHostInterface(ctx, vmPoolId, vmPoolClusterHostId, vmPoolClusterHostInterfaceId)
Retrieves a VM Cluster Host Interface

Returns a VM Cluster Host Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 
  **vmPoolClusterHostInterfaceId** | **float64**|  | 

### Return type

[**VmPoolHostInterfaces**](VMPoolHostInterfaces.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMPoolClusterHostInterfaces**
> []VmPoolHostInterfaces GetVMPoolClusterHostInterfaces(ctx, vmPoolId, vmPoolClusterHostId)
Retrieves a list of VM Cluster Host Interfaces

Returns a list of VM Cluster Host Interfaces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 

### Return type

[**[]VmPoolHostInterfaces**](VMPoolHostInterfaces.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMPoolClusterHostVMs**
> VmList GetVMPoolClusterHostVMs(ctx, vmPoolId, vmPoolClusterHostId)
Retrieves a list of VM Cluster Host VMs

Returns a list of VM Cluster Host VMs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 

### Return type

[**VmList**](VMList.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMPoolClusterHosts**
> VmPoolHostsList GetVMPoolClusterHosts(ctx, vmPoolId)
Get list of VM Cluster Hosts linked to the VM Pool

Returns list of VM Cluster Hosts linked to the VM Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 

### Return type

[**VmPoolHostsList**](VMPoolHostsList.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMPoolVMs**
> GetVMPoolVMs(ctx, vmPoolId)
Returns all VMs linked to the VM Pool

Returns all VMs linked to the VM Pool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmPoolId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVMPools**
> VmPoolList GetVMPools(ctx, )
Get all VM Pools

Returns list of all VM Pools

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**VmPoolList**](VMPoolList.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVMPool**
> VmPool UpdateVMPool(ctx, body, vmPoolId)
Updates VM Pool information

Updates VM Pool information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmPool**](UpdateVmPool.md)| The VM Pool update object | 
  **vmPoolId** | **float64**|  | 

### Return type

[**VmPool**](VMPool.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVMPoolClusterHostInterface**
> VmPoolHostInterfaces UpdateVMPoolClusterHostInterface(ctx, body, vmPoolId, vmPoolClusterHostId, vmPoolClusterHostInterfaceId)
Updates a VM Cluster Host Interface

Updates a VM Cluster Host Interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmPoolClusterHostInterface**](UpdateVmPoolClusterHostInterface.md)| The VM Pool Cluster Host Interface update object | 
  **vmPoolId** | **float64**|  | 
  **vmPoolClusterHostId** | **float64**|  | 
  **vmPoolClusterHostInterfaceId** | **float64**|  | 

### Return type

[**VmPoolHostInterfaces**](VMPoolHostInterfaces.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

