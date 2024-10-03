# {{classname}}

All URIs are relative to **

Method | HTTP request | Description
------------- | ------------- | -------------
[**InventoryController1CreateStorage**](StorageApi.md#InventoryController1CreateStorage) | **Post** /api/v2/storages | Creates a Storage

# **InventoryController1CreateStorage**
> StorageRegistrationResponse InventoryController1CreateStorage(ctx, body)
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

