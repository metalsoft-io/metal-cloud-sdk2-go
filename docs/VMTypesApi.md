# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryController1CreateVMType**](VMTypesApi.md#InventoryController1CreateVMType) | **Post** /api/v2/vm-types | Creates a VM Type
[**InventoryController1DeleteVMType**](VMTypesApi.md#InventoryController1DeleteVMType) | **Delete** /api/v2/vm-types/{vmTypeId} | Deletes a VM Type
[**InventoryController1GetVMType**](VMTypesApi.md#InventoryController1GetVMType) | **Get** /api/v2/vm-types/{vmTypeId} | Get VM Type information
[**InventoryController1GetVMTypes**](VMTypesApi.md#InventoryController1GetVMTypes) | **Get** /api/v2/vm-types | Get all VM Types
[**InventoryController1GetVMsByVMType**](VMTypesApi.md#InventoryController1GetVMsByVMType) | **Get** /api/v2/vm-types/{vmTypeId}/vms | Returns all VMs linked to the VM Type
[**InventoryController1UpdateVMType**](VMTypesApi.md#InventoryController1UpdateVMType) | **Patch** /api/v2/vm-types/{vmTypeId} | Updates VM Type information

# **InventoryController1CreateVMType**
> VmTypeDto InventoryController1CreateVMType(ctx, body)
Creates a VM Type

Creates a VM Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVmTypeDto**](CreateVmTypeDto.md)| The VM Type create object | 

### Return type

[**VmTypeDto**](VMTypeDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1DeleteVMType**
> InventoryController1DeleteVMType(ctx, vmTypeId)
Deletes a VM Type

Deletes a VM Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmTypeId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMType**
> VmTypeDto InventoryController1GetVMType(ctx, vmTypeId)
Get VM Type information

Returns VM Type information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmTypeId** | **float64**|  | 

### Return type

[**VmTypeDto**](VMTypeDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMTypes**
> []VmTypeDto InventoryController1GetVMTypes(ctx, )
Get all VM Types

Returns list of all VM Types

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]VmTypeDto**](VMTypeDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1GetVMsByVMType**
> InventoryController1GetVMsByVMType(ctx, vmTypeId)
Returns all VMs linked to the VM Type

Returns all VMs linked to the VM Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmTypeId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InventoryController1UpdateVMType**
> VmTypeDto InventoryController1UpdateVMType(ctx, body, vmTypeId)
Updates VM Type information

Updates VM Type information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVmTypeDto**](UpdateVmTypeDto.md)| The VM Type update object | 
  **vmTypeId** | **float64**|  | 

### Return type

[**VmTypeDto**](VMTypeDto.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

