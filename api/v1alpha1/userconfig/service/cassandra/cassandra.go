// Code generated by user config generator. DO NOT EDIT.
// +kubebuilder:object:generate=true

package cassandrauserconfig

// cassandra configuration values
type Cassandra struct {
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1000000
	// Fail any multiple-partition batch exceeding this value. 50kb (10x warn threshold) by default.
	BatchSizeFailThresholdInKb *int `groups:"create,update" json:"batch_size_fail_threshold_in_kb,omitempty"`

	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1000000
	// Log a warning message on any multiple-partition batch size exceeding this value.5kb per batch by default.Caution should be taken on increasing the size of this thresholdas it can lead to node instability.
	BatchSizeWarnThresholdInKb *int `groups:"create,update" json:"batch_size_warn_threshold_in_kb,omitempty"`

	// +kubebuilder:validation:MaxLength=128
	// Name of the datacenter to which nodes of this service belong. Can be set only when creating the service.
	Datacenter *string `groups:"create,update" json:"datacenter,omitempty"`
}

// CIDR address block, either as a string, or in a dict with an optional description field
type IpFilter struct {
	// +kubebuilder:validation:MaxLength=1024
	// Description for IP filter list entry
	Description *string `groups:"create,update" json:"description,omitempty"`

	// +kubebuilder:validation:MaxLength=43
	// CIDR address block
	Network string `groups:"create,update" json:"network"`
}

// Allow access to selected service ports from private networks
type PrivateAccess struct {
	// Allow clients to connect to prometheus with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}

// Allow access to selected service ports from the public Internet
type PublicAccess struct {
	// Allow clients to connect to prometheus from the public internet for service nodes that are in a project VPC or another type of private network
	Prometheus *bool `groups:"create,update" json:"prometheus,omitempty"`
}
type CassandraUserConfig struct {
	// +kubebuilder:validation:MaxItems=1
	// Additional Cloud Regions for Backup Replication
	AdditionalBackupRegions []string `groups:"create,update" json:"additional_backup_regions,omitempty"`

	// cassandra configuration values
	Cassandra *Cassandra `groups:"create,update" json:"cassandra,omitempty"`

	// +kubebuilder:validation:Enum="4";"3"
	// Cassandra major version
	CassandraVersion *string `groups:"create,update" json:"cassandra_version,omitempty"`

	// +kubebuilder:validation:MaxItems=1024
	// Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'
	IpFilter []*IpFilter `groups:"create,update" json:"ip_filter,omitempty"`

	// Sets the service into migration mode enabling the sstableloader utility to be used to upload Cassandra data files. Available only on service create.
	MigrateSstableloader *bool `groups:"create,update" json:"migrate_sstableloader,omitempty"`

	// Allow access to selected service ports from private networks
	PrivateAccess *PrivateAccess `groups:"create,update" json:"private_access,omitempty"`

	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Name of another project to fork a service from. This has effect only when a new service is being created.
	ProjectToForkFrom *string `groups:"create" json:"project_to_fork_from,omitempty"`

	// Allow access to selected service ports from the public Internet
	PublicAccess *PublicAccess `groups:"create,update" json:"public_access,omitempty"`

	// +kubebuilder:validation:MaxLength=64
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Value is immutable"
	// Name of another service to fork from. This has effect only when a new service is being created.
	ServiceToForkFrom *string `groups:"create" json:"service_to_fork_from,omitempty"`

	// +kubebuilder:validation:MaxLength=64
	// When bootstrapping, instead of creating a new Cassandra cluster try to join an existing one from another service. Can only be set on service creation.
	ServiceToJoinWith *string `groups:"create,update" json:"service_to_join_with,omitempty"`

	// Use static public IP addresses
	StaticIps *bool `groups:"create,update" json:"static_ips,omitempty"`
}
