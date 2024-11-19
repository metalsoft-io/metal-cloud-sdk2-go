# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReProvisionNetworkEquipment**](NetworkEquipmentApi.md#ReProvisionNetworkEquipment) | **Post** /api/v2/network-equipment/re-provision | Re-provision network equipment

# **ReProvisionNetworkEquipment**
> NetworkEquipmentReProvisionResponse ReProvisionNetworkEquipment(ctx, body)
Re-provision network equipment

Re-provision network equipment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NetworkEquipmentReProvisionOptions**](NetworkEquipmentReProvisionOptions.md)| The network equipment re-provision options | 

### Return type

[**NetworkEquipmentReProvisionResponse**](NetworkEquipmentReProvisionResponse.md)

### Authorization

[JWT](../README.md#JWT), [apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

