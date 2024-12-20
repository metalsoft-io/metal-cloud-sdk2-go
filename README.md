# Go API client for sdk2

MetalSoft REST API documentation

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0
- Package version: 6.4.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://www.metalsoft.io/contact](https://www.metalsoft.io/contact)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./sdk2"
```

## Documentation for API Endpoints

All URIs are relative to **

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AIApi* | [**GenerateEliResponse**](docs/AIApi.md#generateeliresponse) | **Post** /api/v2/ai/generate | Request from AI a response for the given input
*AccountsApi* | [**ArchiveAccount**](docs/AccountsApi.md#archiveaccount) | **Post** /api/v2/accounts/{accountId}/actions/archive | Archive account
*AccountsApi* | [**CreateAccount**](docs/AccountsApi.md#createaccount) | **Post** /api/v2/accounts | Create account
*AccountsApi* | [**GetAccount**](docs/AccountsApi.md#getaccount) | **Get** /api/v2/accounts/{accountId} | Get account by id
*AccountsApi* | [**GetAccountLimits**](docs/AccountsApi.md#getaccountlimits) | **Get** /api/v2/accounts/{accountId}/limits | Get account limits
*AccountsApi* | [**GetAccountUsers**](docs/AccountsApi.md#getaccountusers) | **Get** /api/v2/accounts/{accountId}/users | Get users for account
*AccountsApi* | [**GetAccounts**](docs/AccountsApi.md#getaccounts) | **Get** /api/v2/accounts | Get all accounts
*AccountsApi* | [**UnarchiveAccount**](docs/AccountsApi.md#unarchiveaccount) | **Post** /api/v2/accounts/{accountId}/actions/unarchive | Unarchive account
*AccountsApi* | [**UpdateAccount**](docs/AccountsApi.md#updateaccount) | **Patch** /api/v2/accounts/{accountId} | Update account
*AccountsApi* | [**UpdateAccountLimits**](docs/AccountsApi.md#updateaccountlimits) | **Patch** /api/v2/accounts/{accountId}/limits | Update account limits
*BucketApi* | [**CreateInfrastructureBucket**](docs/BucketApi.md#createinfrastructurebucket) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets | Creates a Bucket
*BucketApi* | [**DeleteBucket**](docs/BucketApi.md#deletebucket) | **Delete** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId} | Deletes a Bucket
*BucketApi* | [**GetBucket**](docs/BucketApi.md#getbucket) | **Get** /api/v2/buckets/{bucketId} | Get Bucket information
*BucketApi* | [**GetInfrastructureBucket**](docs/BucketApi.md#getinfrastructurebucket) | **Get** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId} | Get Bucket information
*BucketApi* | [**GetInfrastructureBuckets**](docs/BucketApi.md#getinfrastructurebuckets) | **Get** /api/v2/infrastructures/{infrastructureId}/buckets | Get all Buckets
*BucketApi* | [**StartBucket**](docs/BucketApi.md#startbucket) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId}/actions/start | Starts a Bucket
*BucketApi* | [**StopBucket**](docs/BucketApi.md#stopbucket) | **Post** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId}/actions/stop | Stops a Bucket
*BucketApi* | [**UpdateBucket**](docs/BucketApi.md#updatebucket) | **Patch** /api/v2/infrastructures/{infrastructureId}/buckets/{bucketId} | Updates Bucket information
*ConfigurationApi* | [**GetConfiguration**](docs/ConfigurationApi.md#getconfiguration) | **Get** /api/v2/config | Get configuration
*ExtensionApi* | [**ArchiveExtension**](docs/ExtensionApi.md#archiveextension) | **Post** /api/v2/extensions/{extensionId}/actions/archive | Archive published extension
*ExtensionApi* | [**CreateExtension**](docs/ExtensionApi.md#createextension) | **Post** /api/v2/extensions | Create extension
*ExtensionApi* | [**GetExtension**](docs/ExtensionApi.md#getextension) | **Get** /api/v2/extensions/{extensionId} | Get details for an extension
*ExtensionApi* | [**GetExtensions**](docs/ExtensionApi.md#getextensions) | **Get** /api/v2/extensions | List available extensions
*ExtensionApi* | [**PublishExtension**](docs/ExtensionApi.md#publishextension) | **Post** /api/v2/extensions/{extensionId}/actions/publish | Publish draft extension
*ExtensionApi* | [**UpdateExtension**](docs/ExtensionApi.md#updateextension) | **Patch** /api/v2/extensions/{extensionId} | Update extension
*ExtensionInstanceApi* | [**CreateExtensionInstance**](docs/ExtensionInstanceApi.md#createextensioninstance) | **Post** /api/v2/infrastructures/{infrastructureId}/extension-instances | Add extension instance to an infrastructure
*ExtensionInstanceApi* | [**DeleteExtensionInstance**](docs/ExtensionInstanceApi.md#deleteextensioninstance) | **Delete** /api/v2/extension-instances/{extensionInstanceId} | Delete extension instance
*ExtensionInstanceApi* | [**GetExtensionInstance**](docs/ExtensionInstanceApi.md#getextensioninstance) | **Get** /api/v2/extension-instances/{extensionInstanceId} | Get extension instance details
*ExtensionInstanceApi* | [**GetExtensionInstances**](docs/ExtensionInstanceApi.md#getextensioninstances) | **Get** /api/v2/extension-instances | Get extension instances list
*ExtensionInstanceApi* | [**UpdateExtensionInstance**](docs/ExtensionInstanceApi.md#updateextensioninstance) | **Patch** /api/v2/extension-instances/{extensionInstanceId} | Update extension instance configuration
*FileShareApi* | [**CreateInfrastructureFileShare**](docs/FileShareApi.md#createinfrastructurefileshare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares | Creates a File Share
*FileShareApi* | [**DeleteFileShare**](docs/FileShareApi.md#deletefileshare) | **Delete** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Deletes a File Share
*FileShareApi* | [**GetFileShare**](docs/FileShareApi.md#getfileshare) | **Get** /api/v2/file-shares/{fileShareId} | Get File Share information
*FileShareApi* | [**GetFileShareHosts**](docs/FileShareApi.md#getfilesharehosts) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/hosts | Get the Hosts of File Share
*FileShareApi* | [**GetInfrastructureFileShare**](docs/FileShareApi.md#getinfrastructurefileshare) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Get File Share information
*FileShareApi* | [**GetInfrastructureFileShares**](docs/FileShareApi.md#getinfrastructurefileshares) | **Get** /api/v2/infrastructures/{infrastructureId}/file-shares | Get all File Shares
*FileShareApi* | [**StartFileShare**](docs/FileShareApi.md#startfileshare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/start | Starts a File Share
*FileShareApi* | [**StopFileShare**](docs/FileShareApi.md#stopfileshare) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/stop | Stops a File Share
*FileShareApi* | [**UpdateFileShare**](docs/FileShareApi.md#updatefileshare) | **Patch** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId} | Updates File Share information
*FileShareApi* | [**UpdateFileShareInstanceArrayHostsBulk**](docs/FileShareApi.md#updatefileshareinstancearrayhostsbulk) | **Post** /api/v2/infrastructures/{infrastructureId}/file-shares/{fileShareId}/actions/modify-instance-array-hosts-bulk | Updates Instance Array Hosts on the File Share
*InfrastructureApi* | [**DeployInfrastructure**](docs/InfrastructureApi.md#deployinfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/deploy | Deploys the specified infrastructure
*InfrastructureApi* | [**GetInfrastructure**](docs/InfrastructureApi.md#getinfrastructure) | **Get** /api/v2/infrastructures/{infrastructureId} | Retrieves the specified infrastructure
*InfrastructureApi* | [**GetInfrastructures**](docs/InfrastructureApi.md#getinfrastructures) | **Get** /api/v2/infrastructures | Get all infrastructures
*InfrastructureApi* | [**RevertInfrastructure**](docs/InfrastructureApi.md#revertinfrastructure) | **Post** /api/v2/infrastructures/{infrastructureId}/actions/revert | Reverts the specified infrastructure
*NetworkApi* | [**CreateInfrastructureNetwork**](docs/NetworkApi.md#createinfrastructurenetwork) | **Post** /api/v2/infrastructures/{infrastructureId}/networks | Creates a new LAN network on the infrastructure
*NetworkApi* | [**DeleteInfrastructureNetwork**](docs/NetworkApi.md#deleteinfrastructurenetwork) | **Delete** /api/v2/infrastructures/{infrastructureId}/networks/{networkId} | Deletes a network from the infrastructure
*NetworkApi* | [**GetInfrastructureNetwork**](docs/NetworkApi.md#getinfrastructurenetwork) | **Get** /api/v2/infrastructures/{infrastructureId}/networks/{networkId} | Gets the specified network from the infrastructure
*NetworkApi* | [**GetInfrastructureNetworks**](docs/NetworkApi.md#getinfrastructurenetworks) | **Get** /api/v2/infrastructures/{infrastructureId}/networks | Retrieves all networks on the infrastructure
*NetworkDevicesApi* | [**ChangeNetworkDeviceStatus**](docs/NetworkDevicesApi.md#changenetworkdevicestatus) | **Patch** /api/v2/network-devices/{networkDeviceId}/actions/change-status | Change status of a network device
*NetworkDevicesApi* | [**DiscoverNetworkDevice**](docs/NetworkDevicesApi.md#discovernetworkdevice) | **Post** /api/v2/network-devices/{networkDeviceId}/discover | Discover network device interfaces, hardware and software configuration
*NetworkDevicesApi* | [**EnableNetworkDeviceSyslog**](docs/NetworkDevicesApi.md#enablenetworkdevicesyslog) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/syslog-subscribe | Enables remote syslog for a network device
*NetworkDevicesApi* | [**GetNetworkDevicePorts**](docs/NetworkDevicesApi.md#getnetworkdeviceports) | **Get** /api/v2/network-devices/{networkDeviceId}/ports | Get all ports for network device
*NetworkDevicesApi* | [**ResetNetworkDevice**](docs/NetworkDevicesApi.md#resetnetworkdevice) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/reset | Resets a network device to default state
*NetworkDevicesApi* | [**SetNetworkDevicePortStatus**](docs/NetworkDevicesApi.md#setnetworkdeviceportstatus) | **Post** /api/v2/network-devices/{networkDeviceId}/actions/set-port-status | Set port status
*NetworkEquipmentApi* | [**ReProvisionNetworkEquipment**](docs/NetworkEquipmentApi.md#reprovisionnetworkequipment) | **Post** /api/v2/network-equipment/re-provision | Re-provision network equipment
*ResourcePoolsApi* | [**AddResourcePoolUser**](docs/ResourcePoolsApi.md#addresourcepooluser) | **Post** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Add a user to a Resource Pool
*ResourcePoolsApi* | [**AddServerToResourcePool**](docs/ResourcePoolsApi.md#addservertoresourcepool) | **Put** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Add a server to a Resource Pool
*ResourcePoolsApi* | [**AddSubnetPoolToResourcePool**](docs/ResourcePoolsApi.md#addsubnetpooltoresourcepool) | **Put** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Add a subnet pool to a resource pool
*ResourcePoolsApi* | [**CreateResourcePool**](docs/ResourcePoolsApi.md#createresourcepool) | **Post** /api/v2/resource-pools | Creates a Resource Pool
*ResourcePoolsApi* | [**DeleteResourcePool**](docs/ResourcePoolsApi.md#deleteresourcepool) | **Delete** /api/v2/resource-pools/{resourcePoolId} | Deletes a Resource Pool
*ResourcePoolsApi* | [**GetResourcePool**](docs/ResourcePoolsApi.md#getresourcepool) | **Get** /api/v2/resource-pools/{resourcePoolId} | Get Resource Pool information
*ResourcePoolsApi* | [**GetResourcePoolServers**](docs/ResourcePoolsApi.md#getresourcepoolservers) | **Get** /api/v2/resource-pools/{resourcePoolId}/servers | Get all servers that are part of a Resource Pool
*ResourcePoolsApi* | [**GetResourcePoolSubnetPools**](docs/ResourcePoolsApi.md#getresourcepoolsubnetpools) | **Get** /api/v2/resource-pools/{resourcePoolId}/subnet-pools | Get all subnet pools that are part of a resource pool
*ResourcePoolsApi* | [**GetResourcePoolUsers**](docs/ResourcePoolsApi.md#getresourcepoolusers) | **Get** /api/v2/resource-pools/{resourcePoolId}/users | Get all users that have access to a Resource Pool
*ResourcePoolsApi* | [**GetResourcePools**](docs/ResourcePoolsApi.md#getresourcepools) | **Get** /api/v2/resource-pools | Get all Resource Pools
*ResourcePoolsApi* | [**GetUserResourcePools**](docs/ResourcePoolsApi.md#getuserresourcepools) | **Get** /api/v2/resource-pools/user/{userId} | Get all Resource Pools that a user has access to
*ResourcePoolsApi* | [**RemoveResourcePoolUser**](docs/ResourcePoolsApi.md#removeresourcepooluser) | **Delete** /api/v2/resource-pools/user/{userId}/pool/{resourcePoolId} | Remove a user from a Resource Pool
*ResourcePoolsApi* | [**RemoveServerFromResourcePool**](docs/ResourcePoolsApi.md#removeserverfromresourcepool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/server/{serverId} | Remove a server from a Resource Pool
*ResourcePoolsApi* | [**RemoveSubnetPoolFromResourcePool**](docs/ResourcePoolsApi.md#removesubnetpoolfromresourcepool) | **Delete** /api/v2/resource-pools/{resourcePoolId}/subnet-pool/{subnetPoolId} | Remove a subnet from a resource pool
*ResourcePoolsApi* | [**UpdateResourcePool**](docs/ResourcePoolsApi.md#updateresourcepool) | **Put** /api/v2/resource-pools/{resourcePoolId} | Updates Resource Pool information
*SecurityApi* | [**ListProviders**](docs/SecurityApi.md#listproviders) | **Get** /api/v2/authentication/providers | Get available authentication providers
*SecurityApi* | [**UpdateProvider**](docs/SecurityApi.md#updateprovider) | **Patch** /api/v2/authentication/providers/{name} | Updates authentication provider
*ServerApi* | [**EnableServerSyslog**](docs/ServerApi.md#enableserversyslog) | **Post** /api/v2/servers/{serverId}/actions/syslog-subscribe | Enables remote syslog for a server
*ServerApi* | [**GetServerInfo**](docs/ServerApi.md#getserverinfo) | **Get** /api/v2/servers/{serverId} | Get Server information
*ServerApi* | [**GetServerPowerState**](docs/ServerApi.md#getserverpowerstate) | **Post** /api/v2/servers/{serverId}/actions/get-power | Gets the power state of a server
*ServerApi* | [**GetServerRemoteConsoleInfo**](docs/ServerApi.md#getserverremoteconsoleinfo) | **Get** /api/v2/servers/{serverId}/remote-console-info | Get Remote Console information
*ServerApi* | [**GetServerVNCInfo**](docs/ServerApi.md#getservervncinfo) | **Get** /api/v2/servers/{serverId}/vnc-info | Get VNC information
*ServerApi* | [**GetServers**](docs/ServerApi.md#getservers) | **Get** /api/v2/servers | Get a list of Servers
*ServerApi* | [**ReRegisterServer**](docs/ServerApi.md#reregisterserver) | **Post** /api/v2/servers/{serverId}/actions/re-register | Re-register a server
*ServerApi* | [**RegisterServer**](docs/ServerApi.md#registerserver) | **Post** /api/v2/servers | Initialize server registration
*ServerApi* | [**ResetServerToFactoryDefaults**](docs/ServerApi.md#resetservertofactorydefaults) | **Post** /api/v2/servers/{serverId}/actions/factory-reset | Resets a server to factory defaults
*ServerApi* | [**SetServerPowerState**](docs/ServerApi.md#setserverpowerstate) | **Post** /api/v2/servers/{serverId}/actions/set-power | Sets the power state of a server
*StorageApi* | [**CreateStorage**](docs/StorageApi.md#createstorage) | **Post** /api/v2/storages | Creates a Storage
*StorageApi* | [**GetStorage**](docs/StorageApi.md#getstorage) | **Get** /api/v2/storages/{storageId} | Retrieves a Storage
*StorageApi* | [**GetStorageBuckets**](docs/StorageApi.md#getstoragebuckets) | **Get** /api/v2/storages/{storageId}/buckets | Get all Buckets linked to the specified storage
*StorageApi* | [**GetStorageFileShares**](docs/StorageApi.md#getstoragefileshares) | **Get** /api/v2/storages/{storageId}/file-shares | Get all File Shares linked to the specified storage
*StorageApi* | [**GetStorages**](docs/StorageApi.md#getstorages) | **Get** /api/v2/storages | Get a list of Storages
*StorageApi* | [**UpdateStorage**](docs/StorageApi.md#updatestorage) | **Patch** /api/v2/storages/{storageId} | Updates a Storage
*SystemApi* | [**GetVersion**](docs/SystemApi.md#getversion) | **Get** /api/v2/version | Get MetalSoft system version
*UsersApi* | [**ArchiveUser**](docs/UsersApi.md#archiveuser) | **Post** /api/v2/users/{userId}/actions/archive | Archive user
*UsersApi* | [**ChangeUserAccount**](docs/UsersApi.md#changeuseraccount) | **Post** /api/v2/users/{userId}/actions/change-account | Change account for user
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /api/v2/users | Creates a user
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /api/v2/users/{userId} | Get user
*UsersApi* | [**GetUserLimits**](docs/UsersApi.md#getuserlimits) | **Get** /api/v2/users/{userId}/limits | Get user limits
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /api/v2/users | Get users
*UsersApi* | [**UnarchiveUser**](docs/UsersApi.md#unarchiveuser) | **Post** /api/v2/users/{userId}/actions/unarchive | Unarchive user
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Patch** /api/v2/users/{userId} | Update user
*UsersApi* | [**UpdateUserLimits**](docs/UsersApi.md#updateuserlimits) | **Patch** /api/v2/users/{userId}/limits | Update user limits
*VMInstanceApi* | [**ApplyVMTypeOnVMInstance**](docs/VMInstanceApi.md#applyvmtypeonvminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance
*VMInstanceApi* | [**CreateVMInstance**](docs/VMInstanceApi.md#createvminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances | Creates a VM Instance
*VMInstanceApi* | [**DeleteVMInstance**](docs/VMInstanceApi.md#deletevminstance) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Deletes a VM Instance
*VMInstanceApi* | [**GetVMInstance**](docs/VMInstanceApi.md#getvminstance) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Get VM Instance information
*VMInstanceApi* | [**GetVMInstancePowerStatus**](docs/VMInstanceApi.md#getvminstancepowerstatus) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/power-status | Retrieves the power status of the VM Instance
*VMInstanceApi* | [**RebootVMInstance**](docs/VMInstanceApi.md#rebootvminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/reboot | Reboots the VM Instance
*VMInstanceApi* | [**ShutdownVMInstance**](docs/VMInstanceApi.md#shutdownvminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/shutdown | Shuts down the VM Instance
*VMInstanceApi* | [**StartVMInstance**](docs/VMInstanceApi.md#startvminstance) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId}/start | Starts the VM Instance
*VMInstanceApi* | [**UpdateVMInstance**](docs/VMInstanceApi.md#updatevminstance) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instances/{vmInstanceId} | Updates VM Instance information
*VMInstanceGroupApi* | [**ApplyVMTypeOnVMInstanceGroup**](docs/VMInstanceGroupApi.md#applyvmtypeonvminstancegroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/actions/apply-type/{vmTypeId} | Applies a VM Type to a VM Instance Group
*VMInstanceGroupApi* | [**CreateVMInstanceGroup**](docs/VMInstanceGroupApi.md#createvminstancegroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Creates a VM Instance Group
*VMInstanceGroupApi* | [**CreateVMInterfaceOnVMInstanceGroup**](docs/VMInstanceGroupApi.md#createvminterfaceonvminstancegroup) | **Post** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/interfaces | Creates a new Virtual Interface for the VM Instance Group
*VMInstanceGroupApi* | [**DeleteVMInstanceGroup**](docs/VMInstanceGroupApi.md#deletevminstancegroup) | **Delete** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Deletes a VM Instance Group
*VMInstanceGroupApi* | [**GetVMInstanceGroup**](docs/VMInstanceGroupApi.md#getvminstancegroup) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Get VM Instance Group information
*VMInstanceGroupApi* | [**GetVMInstanceGroupVMInstances**](docs/VMInstanceGroupApi.md#getvminstancegroupvminstances) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/vm-instances | Get the VM Instances of VM Instance Group
*VMInstanceGroupApi* | [**GetVMInstanceGroups**](docs/VMInstanceGroupApi.md#getvminstancegroups) | **Get** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups | Get all VM Instance Groups
*VMInstanceGroupApi* | [**UpdateNetworkProfileOnVMInstanceGroupNetwork**](docs/VMInstanceGroupApi.md#updatenetworkprofileonvminstancegroupnetwork) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId}/network/{networkId} | Applies the given Network Profile to the specified VM Instance Group Network
*VMInstanceGroupApi* | [**UpdateVMInstanceGroup**](docs/VMInstanceGroupApi.md#updatevminstancegroup) | **Patch** /api/v2/infrastructures/{infrastructureId}/vm-instance-groups/{vmInstanceGroupId} | Updates VM Instance Group information
*VMPoolsApi* | [**CreateVMPool**](docs/VMPoolsApi.md#createvmpool) | **Post** /api/v2/vm-pools | Creates a VM Pool
*VMPoolsApi* | [**DeleteVMPool**](docs/VMPoolsApi.md#deletevmpool) | **Delete** /api/v2/vm-pools/{vmPoolId} | Deletes a VM Pool
*VMPoolsApi* | [**GetVMPool**](docs/VMPoolsApi.md#getvmpool) | **Get** /api/v2/vm-pools/{vmPoolId} | Get VM Pool information
*VMPoolsApi* | [**GetVMPoolClusterHost**](docs/VMPoolsApi.md#getvmpoolclusterhost) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId} | Retrieves a VM Cluster Host
*VMPoolsApi* | [**GetVMPoolClusterHostInterface**](docs/VMPoolsApi.md#getvmpoolclusterhostinterface) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Retrieves a VM Cluster Host Interface
*VMPoolsApi* | [**GetVMPoolClusterHostInterfaces**](docs/VMPoolsApi.md#getvmpoolclusterhostinterfaces) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces | Retrieves a list of VM Cluster Host Interfaces
*VMPoolsApi* | [**GetVMPoolClusterHostVMs**](docs/VMPoolsApi.md#getvmpoolclusterhostvms) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/vms | Retrieves a list of VM Cluster Host VMs
*VMPoolsApi* | [**GetVMPoolClusterHosts**](docs/VMPoolsApi.md#getvmpoolclusterhosts) | **Get** /api/v2/vm-pools/{vmPoolId}/cluster-hosts | Get list of VM Cluster Hosts linked to the VM Pool
*VMPoolsApi* | [**GetVMPoolVMs**](docs/VMPoolsApi.md#getvmpoolvms) | **Get** /api/v2/vm-pools/{vmPoolId}/vms | Returns all VMs linked to the VM Pool
*VMPoolsApi* | [**GetVMPools**](docs/VMPoolsApi.md#getvmpools) | **Get** /api/v2/vm-pools | Get all VM Pools
*VMPoolsApi* | [**UpdateVMPool**](docs/VMPoolsApi.md#updatevmpool) | **Patch** /api/v2/vm-pools/{vmPoolId} | Updates VM Pool information
*VMPoolsApi* | [**UpdateVMPoolClusterHostInterface**](docs/VMPoolsApi.md#updatevmpoolclusterhostinterface) | **Patch** /api/v2/vm-pools/{vmPoolId}/cluster-hosts/{vmPoolClusterHostId}/interfaces/{vmPoolClusterHostInterfaceId} | Updates a VM Cluster Host Interface
*VMTypesApi* | [**CreateVMType**](docs/VMTypesApi.md#createvmtype) | **Post** /api/v2/vm-types | Creates a VM Type
*VMTypesApi* | [**DeleteVMType**](docs/VMTypesApi.md#deletevmtype) | **Delete** /api/v2/vm-types/{vmTypeId} | Deletes a VM Type
*VMTypesApi* | [**GetVMType**](docs/VMTypesApi.md#getvmtype) | **Get** /api/v2/vm-types/{vmTypeId} | Get VM Type information
*VMTypesApi* | [**GetVMTypes**](docs/VMTypesApi.md#getvmtypes) | **Get** /api/v2/vm-types | Get all VM Types
*VMTypesApi* | [**GetVMsByVMType**](docs/VMTypesApi.md#getvmsbyvmtype) | **Get** /api/v2/vm-types/{vmTypeId}/vms | Returns all VMs linked to the VM Type
*VMTypesApi* | [**UpdateVMType**](docs/VMTypesApi.md#updatevmtype) | **Patch** /api/v2/vm-types/{vmTypeId} | Updates VM Type information
*VMsApi* | [**GetVM**](docs/VMsApi.md#getvm) | **Get** /api/v2/vms/{vmId} | Retrieves the VM information
*VMsApi* | [**GetVMPowerStatus**](docs/VMsApi.md#getvmpowerstatus) | **Get** /api/v2/vms/{vmId}/power-status | Retrieves the power status of the VM
*VMsApi* | [**GetVMRemoteConsoleInfo**](docs/VMsApi.md#getvmremoteconsoleinfo) | **Get** /api/v2/vms/{vmId}/remote-console-info | Get Remote Console information
*VMsApi* | [**RebootVM**](docs/VMsApi.md#rebootvm) | **Post** /api/v2/vms/{vmId}/reboot | Reboots the VM
*VMsApi* | [**ShutdownVM**](docs/VMsApi.md#shutdownvm) | **Post** /api/v2/vms/{vmId}/shutdown | Shuts down the VM
*VMsApi* | [**StartVM**](docs/VMsApi.md#startvm) | **Post** /api/v2/vms/{vmId}/start | Starts the VM
*VMsApi* | [**UpdateVM**](docs/VMsApi.md#updatevm) | **Patch** /api/v2/vms/{vmId} | Updates VM information

## Documentation For Models

 - [AccountAddressDto](docs/AccountAddressDto.md)
 - [AccountDto](docs/AccountDto.md)
 - [AccountLimits](docs/AccountLimits.md)
 - [AccountLimitsDto](docs/AccountLimitsDto.md)
 - [AfcJobInfo](docs/AfcJobInfo.md)
 - [AiGenerateRequestDto](docs/AiGenerateRequestDto.md)
 - [AiGenerateResponseDto](docs/AiGenerateResponseDto.md)
 - [AuthenticationProvider](docs/AuthenticationProvider.md)
 - [AuthenticationProviderUpdate](docs/AuthenticationProviderUpdate.md)
 - [Bucket](docs/Bucket.md)
 - [BucketExtendedInfo](docs/BucketExtendedInfo.md)
 - [BucketList](docs/BucketList.md)
 - [ChangeUserAccountDto](docs/ChangeUserAccountDto.md)
 - [CreateAccountDto](docs/CreateAccountDto.md)
 - [CreateBucket](docs/CreateBucket.md)
 - [CreateExtensionDto](docs/CreateExtensionDto.md)
 - [CreateExtensionInstanceDto](docs/CreateExtensionInstanceDto.md)
 - [CreateFileShare](docs/CreateFileShare.md)
 - [CreateNetwork](docs/CreateNetwork.md)
 - [CreateResourcePoolDto](docs/CreateResourcePoolDto.md)
 - [CreateStorage](docs/CreateStorage.md)
 - [CreateUserDto](docs/CreateUserDto.md)
 - [CreateVmInstance](docs/CreateVmInstance.md)
 - [CreateVmInstanceGroup](docs/CreateVmInstanceGroup.md)
 - [CreateVmInstanceGroupInterface](docs/CreateVmInstanceGroupInterface.md)
 - [CreateVmPool](docs/CreateVmPool.md)
 - [CreateVmType](docs/CreateVmType.md)
 - [CustomVariable](docs/CustomVariable.md)
 - [DiscoveryQueryDto](docs/DiscoveryQueryDto.md)
 - [ExtensionActions](docs/ExtensionActions.md)
 - [ExtensionAsset](docs/ExtensionAsset.md)
 - [ExtensionDefinitionDto](docs/ExtensionDefinitionDto.md)
 - [ExtensionDependency](docs/ExtensionDependency.md)
 - [ExtensionDto](docs/ExtensionDto.md)
 - [ExtensionInfoDto](docs/ExtensionInfoDto.md)
 - [ExtensionInfrastructure](docs/ExtensionInfrastructure.md)
 - [ExtensionInput](docs/ExtensionInput.md)
 - [ExtensionInputBoolean](docs/ExtensionInputBoolean.md)
 - [ExtensionInputInteger](docs/ExtensionInputInteger.md)
 - [ExtensionInputOsTemplate](docs/ExtensionInputOsTemplate.md)
 - [ExtensionInputServerType](docs/ExtensionInputServerType.md)
 - [ExtensionInputString](docs/ExtensionInputString.md)
 - [ExtensionInstanceArray](docs/ExtensionInstanceArray.md)
 - [ExtensionInstanceDto](docs/ExtensionInstanceDto.md)
 - [ExtensionInstanceListDto](docs/ExtensionInstanceListDto.md)
 - [ExtensionListDto](docs/ExtensionListDto.md)
 - [ExtensionOutput](docs/ExtensionOutput.md)
 - [ExtensionSharedDrive](docs/ExtensionSharedDrive.md)
 - [ExtensionTask](docs/ExtensionTask.md)
 - [ExtensionTaskAnsible](docs/ExtensionTaskAnsible.md)
 - [ExtensionTaskWebhook](docs/ExtensionTaskWebhook.md)
 - [ExtensionVariable](docs/ExtensionVariable.md)
 - [FileShare](docs/FileShare.md)
 - [FileShareExtendedInfo](docs/FileShareExtendedInfo.md)
 - [FileShareHostBulkOperation](docs/FileShareHostBulkOperation.md)
 - [FileShareHosts](docs/FileShareHosts.md)
 - [FileShareHostsModifyBulk](docs/FileShareHostsModifyBulk.md)
 - [FileShareList](docs/FileShareList.md)
 - [GenericGuiSettings](docs/GenericGuiSettings.md)
 - [InfrastructureDeployOptions](docs/InfrastructureDeployOptions.md)
 - [InfrastructureDeployShutdownOptions](docs/InfrastructureDeployShutdownOptions.md)
 - [InfrastructureDto](docs/InfrastructureDto.md)
 - [Link](docs/Link.md)
 - [Network](docs/Network.md)
 - [NetworkDevicePortStatusDto](docs/NetworkDevicePortStatusDto.md)
 - [NetworkDeviceStatusDto](docs/NetworkDeviceStatusDto.md)
 - [NetworkEquipmentReProvisionOptions](docs/NetworkEquipmentReProvisionOptions.md)
 - [NetworkEquipmentReProvisionResponse](docs/NetworkEquipmentReProvisionResponse.md)
 - [OneOfExtensionInputOptions](docs/OneOfExtensionInputOptions.md)
 - [OneOfExtensionTaskOptions](docs/OneOfExtensionTaskOptions.md)
 - [PatchStorage](docs/PatchStorage.md)
 - [RemoteConsoleInfoDto](docs/RemoteConsoleInfoDto.md)
 - [ResourcePoolDto](docs/ResourcePoolDto.md)
 - [ResourcePoolStatistics](docs/ResourcePoolStatistics.md)
 - [ResourcePoolWithStatsDto](docs/ResourcePoolWithStatsDto.md)
 - [Server](docs/Server.md)
 - [ServerList](docs/ServerList.md)
 - [ServerPowerSetDto](docs/ServerPowerSetDto.md)
 - [ServerReRegistrationResponseDto](docs/ServerReRegistrationResponseDto.md)
 - [ServerRegistrationDto](docs/ServerRegistrationDto.md)
 - [ServerRegistrationResponseDto](docs/ServerRegistrationResponseDto.md)
 - [ServerVncInfoDto](docs/ServerVncInfoDto.md)
 - [Storage](docs/Storage.md)
 - [StorageList](docs/StorageList.md)
 - [StorageRegistrationResponse](docs/StorageRegistrationResponse.md)
 - [StorageRegistrationResponseJobInfo](docs/StorageRegistrationResponseJobInfo.md)
 - [UpdateAccountDto](docs/UpdateAccountDto.md)
 - [UpdateBucket](docs/UpdateBucket.md)
 - [UpdateExtensionDto](docs/UpdateExtensionDto.md)
 - [UpdateExtensionInstanceDto](docs/UpdateExtensionInstanceDto.md)
 - [UpdateFileShare](docs/UpdateFileShare.md)
 - [UpdateResourcePoolDto](docs/UpdateResourcePoolDto.md)
 - [UpdateUserDto](docs/UpdateUserDto.md)
 - [UpdateVm](docs/UpdateVm.md)
 - [UpdateVmInstance](docs/UpdateVmInstance.md)
 - [UpdateVmInstanceGroup](docs/UpdateVmInstanceGroup.md)
 - [UpdateVmInstanceGroupInterface](docs/UpdateVmInstanceGroupInterface.md)
 - [UpdateVmInstanceGroupNetwork](docs/UpdateVmInstanceGroupNetwork.md)
 - [UpdateVmPool](docs/UpdateVmPool.md)
 - [UpdateVmPoolClusterHostInterface](docs/UpdateVmPoolClusterHostInterface.md)
 - [UpdateVmType](docs/UpdateVmType.md)
 - [UserDelegateDto](docs/UserDelegateDto.md)
 - [UserDto](docs/UserDto.md)
 - [UserExperimentalTagDto](docs/UserExperimentalTagDto.md)
 - [UserLimits](docs/UserLimits.md)
 - [UserLimitsDto](docs/UserLimitsDto.md)
 - [UserPromotionDto](docs/UserPromotionDto.md)
 - [UserResourcePermissionDto](docs/UserResourcePermissionDto.md)
 - [UserResourcePermissionsDto](docs/UserResourcePermissionsDto.md)
 - [UserSuspendDto](docs/UserSuspendDto.md)
 - [UserUpdatePasswordDto](docs/UserUpdatePasswordDto.md)
 - [Version](docs/Version.md)
 - [Vm](docs/Vm.md)
 - [VmDisk](docs/VmDisk.md)
 - [VmInstance](docs/VmInstance.md)
 - [VmInstanceGroup](docs/VmInstanceGroup.md)
 - [VmInstanceGroupGuiSettingsDto](docs/VmInstanceGroupGuiSettingsDto.md)
 - [VmInstanceGroupInterface](docs/VmInstanceGroupInterface.md)
 - [VmList](docs/VmList.md)
 - [VmPool](docs/VmPool.md)
 - [VmPoolHostInterfaces](docs/VmPoolHostInterfaces.md)
 - [VmPoolHosts](docs/VmPoolHosts.md)
 - [VmPoolHostsList](docs/VmPoolHostsList.md)
 - [VmPoolList](docs/VmPoolList.md)
 - [VmType](docs/VmType.md)
 - [VmTypeList](docs/VmTypeList.md)

## Documentation For Authorization

## JWT
## apiKey

## Author

support@metalsoft.io
