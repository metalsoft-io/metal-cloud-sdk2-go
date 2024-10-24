/*
 * MetalSoft REST API
 *
 * MetalSoft REST API documentation
 *
 * API version: 2.0
 * Contact: support@metalsoft.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package sdk2

// The limits of the user
type AllOfUserDtoLimits struct {
	ContainerArrayContainersMaxCount float64 `json:"container_array_containers_max_count" yaml:"container_array_containers_max_count"`
	ContainerArrayContainersMinCount float64 `json:"container_array_containers_min_count" yaml:"container_array_containers_min_count"`
	ContainerArrayDriveArraysMaxCount float64 `json:"container_array_drive_arrays_max_count" yaml:"container_array_drive_arrays_max_count"`
	ContainerArrayDriveArraysMinCount float64 `json:"container_array_drive_arrays_min_count" yaml:"container_array_drive_arrays_min_count"`
	ContainerArraySecretsMaxCount float64 `json:"container_array_secrets_max_count" yaml:"container_array_secrets_max_count"`
	ContainerPlatformContainerArrayMaxCount float64 `json:"container_platform_container_array_max_count" yaml:"container_platform_container_array_max_count"`
	DriveArrayDrivesMaxCount float64 `json:"drive_array_drives_max_count" yaml:"drive_array_drives_max_count"`
	DriveArrayDrivesMinCount float64 `json:"drive_array_drives_min_count" yaml:"drive_array_drives_min_count"`
	DriveMaxSizeMbytes float64 `json:"drive_max_size_mbytes" yaml:"drive_max_size_mbytes"`
	DriveMinSizeMbytes float64 `json:"drive_min_size_mbytes" yaml:"drive_min_size_mbytes"`
	InfrastructureActiveMaxCount float64 `json:"infrastructure_active_max_count" yaml:"infrastructure_active_max_count"`
	InfrastructureClusterMaxCount float64 `json:"infrastructure_cluster_max_count" yaml:"infrastructure_cluster_max_count"`
	InfrastructureClusterMysqlAllowed bool `json:"infrastructure_cluster_mysql_allowed" yaml:"infrastructure_cluster_mysql_allowed"`
	InfrastructureContainerClusterKafkaAllowed bool `json:"infrastructure_container_cluster_kafka_allowed" yaml:"infrastructure_container_cluster_kafka_allowed"`
	InfrastructureContainerClusterMaxCount float64 `json:"infrastructure_container_cluster_max_count" yaml:"infrastructure_container_cluster_max_count"`
	InfrastructureContainerClusterPostgresqlAllowed bool `json:"infrastructure_container_cluster_postgresql_allowed" yaml:"infrastructure_container_cluster_postgresql_allowed"`
	InfrastructureContainerClusterSparkArrayAllowed bool `json:"infrastructure_container_cluster_spark_array_allowed" yaml:"infrastructure_container_cluster_spark_array_allowed"`
	InfrastructureContainerClusterSparksqlAllowed bool `json:"infrastructure_container_cluster_sparksql_allowed" yaml:"infrastructure_container_cluster_sparksql_allowed"`
	InfrastructureContainerClusterStreamsetsAllowed bool `json:"infrastructure_container_cluster_streamsets_allowed" yaml:"infrastructure_container_cluster_streamsets_allowed"`
	InfrastructureContainerClusterZookeeperAllowed bool `json:"infrastructure_container_cluster_zookeeper_allowed" yaml:"infrastructure_container_cluster_zookeeper_allowed"`
	InfrastructureContainerClusterZoomdataAllowed bool `json:"infrastructure_container_cluster_zoomdata_allowed" yaml:"infrastructure_container_cluster_zoomdata_allowed"`
	InfrastructureContainerPlatformMaxCount float64 `json:"infrastructure_container_platform_max_count" yaml:"infrastructure_container_platform_max_count"`
	InfrastructureDataLakeEnabled bool `json:"infrastructure_data_lake_enabled" yaml:"infrastructure_data_lake_enabled"`
	InfrastructureDataLakeMaxCount float64 `json:"infrastructure_data_lake_max_count" yaml:"infrastructure_data_lake_max_count"`
	InfrastructureDeletedMaxCount float64 `json:"infrastructure_deleted_max_count" yaml:"infrastructure_deleted_max_count"`
	InfrastructureDriveArrayMaxCount float64 `json:"infrastructure_drive_array_max_count" yaml:"infrastructure_drive_array_max_count"`
	InfrastructureInactiveMaxCount float64 `json:"infrastructure_inactive_max_count" yaml:"infrastructure_inactive_max_count"`
	InfrastructureInstanceArrayMaxCount float64 `json:"infrastructure_instance_array_max_count" yaml:"infrastructure_instance_array_max_count"`
	InfrastructureLanMaxCount float64 `json:"infrastructure_lan_max_count" yaml:"infrastructure_lan_max_count"`
	InfrastructureSanMaxCount float64 `json:"infrastructure_san_max_count" yaml:"infrastructure_san_max_count"`
	InfrastructureSharedDriveMaxCount float64 `json:"infrastructure_shared_drive_max_count" yaml:"infrastructure_shared_drive_max_count"`
	InfrastructureVolumeTemplateExperimentalAllowed bool `json:"infrastructure_volume_template_experimental_allowed" yaml:"infrastructure_volume_template_experimental_allowed"`
	InfrastructureWanMaxCount float64 `json:"infrastructure_wan_max_count" yaml:"infrastructure_wan_max_count"`
	InstanceArrayInstancesMaxCount float64 `json:"instance_array_instances_max_count" yaml:"instance_array_instances_max_count"`
	InstanceArrayInstancesMinCount float64 `json:"instance_array_instances_min_count" yaml:"instance_array_instances_min_count"`
	InfrastructureVmInstanceGroupMaxCount float64 `json:"infrastructure_vm_instance_group_max_count" yaml:"infrastructure_vm_instance_group_max_count"`
	VmInstanceGroupVmInstancesMaxCount float64 `json:"vm_instance_group_vm_instances_max_count" yaml:"vm_instance_group_vm_instances_max_count"`
	VmInstanceMaxDiskSizeMbytes float64 `json:"vm_instance_max_disk_size_mbytes" yaml:"vm_instance_max_disk_size_mbytes"`
	OwnerIsBillable bool `json:"owner_is_billable" yaml:"owner_is_billable"`
	ServerTypeReservationMaxCount float64 `json:"server_type_reservation_max_count" yaml:"server_type_reservation_max_count"`
	ServerTypeReservationMaxQuantity float64 `json:"server_type_reservation_max_quantity" yaml:"server_type_reservation_max_quantity"`
	SharedDriveMaxSizeMbytes float64 `json:"shared_drive_max_size_mbytes" yaml:"shared_drive_max_size_mbytes"`
	SharedDriveMinSizeMbytes float64 `json:"shared_drive_min_size_mbytes" yaml:"shared_drive_min_size_mbytes"`
	InfrastructureFileShareMaxCount float64 `json:"infrastructure_file_share_max_count" yaml:"infrastructure_file_share_max_count"`
	FileShareMinSizeGb float64 `json:"file_share_min_size_gb" yaml:"file_share_min_size_gb"`
	FileShareMaxSizeGb float64 `json:"file_share_max_size_gb" yaml:"file_share_max_size_gb"`
	InfrastructureBucketMaxCount float64 `json:"infrastructure_bucket_max_count" yaml:"infrastructure_bucket_max_count"`
	BucketMinSizeGb float64 `json:"bucket_min_size_gb" yaml:"bucket_min_size_gb"`
	BucketMaxSizeGb float64 `json:"bucket_max_size_gb" yaml:"bucket_max_size_gb"`
	AllowVlanOverrides bool `json:"allow_vlan_overrides" yaml:"allow_vlan_overrides"`
	AllowNetworkProfiles bool `json:"allow_network_profiles" yaml:"allow_network_profiles"`
	ShowOperatingSystemImagesTab bool `json:"show_operating_system_images_tab" yaml:"show_operating_system_images_tab"`
	ShowTemplateAssetsView bool `json:"show_template_assets_view" yaml:"show_template_assets_view"`
	StorageTypes *interface{} `json:"storage_types" yaml:"storage_types"`
	ThresholdMaxCount float64 `json:"threshold_max_count" yaml:"threshold_max_count"`
	UserResourceIscsiStorageSpaceMaxGbytes float64 `json:"user_resource_iscsi_storage_space_max_gbytes" yaml:"user_resource_iscsi_storage_space_max_gbytes"`
	UserResourceServersMaxCount float64 `json:"user_resource_servers_max_count" yaml:"user_resource_servers_max_count"`
	UserResourceServerTypeNameToMaxCount *interface{} `json:"user_resource_server_type_name_to_max_count" yaml:"user_resource_server_type_name_to_max_count"`
	UserSshKeysCountMax float64 `json:"user_ssh_keys_count_max" yaml:"user_ssh_keys_count_max"`
	WanSubnetIpv4MaxCount float64 `json:"wan_subnet_ipv4_max_count" yaml:"wan_subnet_ipv4_max_count"`
	WanSubnetIpv6MaxCount float64 `json:"wan_subnet_ipv6_max_count" yaml:"wan_subnet_ipv6_max_count"`
	WanSubnetPrefixSizeToMaxCount *interface{} `json:"wan_subnet_prefix_size_to_max_count" yaml:"wan_subnet_prefix_size_to_max_count"`
	ShowLegacyPages bool `json:"show_legacy_pages" yaml:"show_legacy_pages"`
	ShowExperimentalPages bool `json:"show_experimental_pages" yaml:"show_experimental_pages"`
	ShowDiagramAppsGlobal bool `json:"show_diagram_apps_global" yaml:"show_diagram_apps_global"`
	ShowDiagramAppClusterTypeCloudera bool `json:"show_diagram_app_cluster_type_cloudera" yaml:"show_diagram_app_cluster_type_cloudera"`
	ShowDiagramAppClusterTypeCouchbase bool `json:"show_diagram_app_cluster_type_couchbase" yaml:"show_diagram_app_cluster_type_couchbase"`
	ShowDiagramAppClusterTypeDatameer bool `json:"show_diagram_app_cluster_type_datameer" yaml:"show_diagram_app_cluster_type_datameer"`
	ShowDiagramAppClusterTypeDatastax bool `json:"show_diagram_app_cluster_type_datastax" yaml:"show_diagram_app_cluster_type_datastax"`
	ShowDiagramAppClusterTypeElasticsearch bool `json:"show_diagram_app_cluster_type_elasticsearch" yaml:"show_diagram_app_cluster_type_elasticsearch"`
	ShowDiagramAppClusterTypeExasol bool `json:"show_diagram_app_cluster_type_exasol" yaml:"show_diagram_app_cluster_type_exasol"`
	ShowDiagramAppClusterTypeHortonworks bool `json:"show_diagram_app_cluster_type_hortonworks" yaml:"show_diagram_app_cluster_type_hortonworks"`
	ShowDiagramAppClusterTypeKubernetes bool `json:"show_diagram_app_cluster_type_kubernetes" yaml:"show_diagram_app_cluster_type_kubernetes"`
	ShowDiagramAppClusterTypeMapr bool `json:"show_diagram_app_cluster_type_mapr" yaml:"show_diagram_app_cluster_type_mapr"`
	ShowDiagramAppClusterTypeMesos bool `json:"show_diagram_app_cluster_type_mesos" yaml:"show_diagram_app_cluster_type_mesos"`
	ShowDiagramAppClusterTypeMysqlPercona bool `json:"show_diagram_app_cluster_type_mysql_percona" yaml:"show_diagram_app_cluster_type_mysql_percona"`
	ShowDiagramAppClusterTypeSplunk bool `json:"show_diagram_app_cluster_type_splunk" yaml:"show_diagram_app_cluster_type_splunk"`
	ShowDiagramAppClusterTypeVmwareVsphere bool `json:"show_diagram_app_cluster_type_vmware_vsphere" yaml:"show_diagram_app_cluster_type_vmware_vsphere"`
	ShowDiagramAppClusterTypeVmwareVcf bool `json:"show_diagram_app_cluster_type_vmware_vcf" yaml:"show_diagram_app_cluster_type_vmware_vcf"`
	ShowDiagramAppClusterTypeKubernetesEksa bool `json:"show_diagram_app_cluster_type_kubernetes_eksa" yaml:"show_diagram_app_cluster_type_kubernetes_eksa"`
	ShowDiagramVmInstanceGroups bool `json:"show_diagram_vm_instance_groups" yaml:"show_diagram_vm_instance_groups"`
	ShowEliChatBot bool `json:"show_eli_chat_bot" yaml:"show_eli_chat_bot"`
	EnableCustomRaidConfiguration bool `json:"enable_custom_raid_configuration" yaml:"enable_custom_raid_configuration"`
	EnableCustomSubnets bool `json:"enable_custom_subnets" yaml:"enable_custom_subnets"`
	ShowStackTrace bool `json:"show_stack_trace" yaml:"show_stack_trace"`
	AllowedNetworkProfiles *interface{} `json:"allowed_network_profiles" yaml:"allowed_network_profiles"`
}
