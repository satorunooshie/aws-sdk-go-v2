// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// An Amazon EKS add-on. For more information, see Amazon EKS add-ons
// (https://docs.aws.amazon.com/eks/latest/userguide/eks-add-ons.html) in the
// Amazon EKS User Guide.
type Addon struct {

	// The Amazon Resource Name (ARN) of the add-on.
	AddonArn *string

	// The name of the add-on.
	AddonName *string

	// The version of the add-on.
	AddonVersion *string

	// The name of the cluster.
	ClusterName *string

	// The date and time that the add-on was created.
	CreatedAt *time.Time

	// An object that represents the health of the add-on.
	Health *AddonHealth

	// The date and time that the add-on was last modified.
	ModifiedAt *time.Time

	// The Amazon Resource Name (ARN) of the IAM role that is bound to the Kubernetes
	// service account used by the add-on.
	ServiceAccountRoleArn *string

	// The status of the add-on.
	Status AddonStatus

	// The metadata that you apply to the add-on to assist with categorization and
	// organization. Each tag consists of a key and an optional value. You define both.
	// Add-on tags do not propagate to any other resources associated with the cluster.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The health of the add-on.
type AddonHealth struct {

	// An object that represents the add-on's health issues.
	Issues []AddonIssue

	noSmithyDocumentSerde
}

// Information about an add-on.
type AddonInfo struct {

	// The name of the add-on.
	AddonName *string

	// An object that represents information about available add-on versions and
	// compatible Kubernetes versions.
	AddonVersions []AddonVersionInfo

	// The type of the add-on.
	Type *string

	noSmithyDocumentSerde
}

// An issue related to an add-on.
type AddonIssue struct {

	// A code that describes the type of issue.
	Code AddonIssueCode

	// A message that provides details about the issue and what might cause it.
	Message *string

	// The resource IDs of the issue.
	ResourceIds []string

	noSmithyDocumentSerde
}

// Information about an add-on version.
type AddonVersionInfo struct {

	// The version of the add-on.
	AddonVersion *string

	// The architectures that the version supports.
	Architecture []string

	// An object that represents the compatibilities of a version.
	Compatibilities []Compatibility

	noSmithyDocumentSerde
}

// An Auto Scaling group that is associated with an Amazon EKS managed node group.
type AutoScalingGroup struct {

	// The name of the Auto Scaling group associated with an Amazon EKS managed node
	// group.
	Name *string

	noSmithyDocumentSerde
}

// An object representing the certificate-authority-data for your cluster.
type Certificate struct {

	// The Base64-encoded certificate data required to communicate with your cluster.
	// Add this to the certificate-authority-data section of the kubeconfig file for
	// your cluster.
	Data *string

	noSmithyDocumentSerde
}

// An object representing an Amazon EKS cluster.
type Cluster struct {

	// The Amazon Resource Name (ARN) of the cluster.
	Arn *string

	// The certificate-authority-data for your cluster.
	CertificateAuthority *Certificate

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The configuration used to connect to a cluster for registration.
	ConnectorConfig *ConnectorConfigResponse

	// The Unix epoch timestamp in seconds for when the cluster was created.
	CreatedAt *time.Time

	// The encryption configuration for the cluster.
	EncryptionConfig []EncryptionConfig

	// The endpoint for your Kubernetes API server.
	Endpoint *string

	// The identity provider information for the cluster.
	Identity *Identity

	// The Kubernetes network configuration for the cluster.
	KubernetesNetworkConfig *KubernetesNetworkConfigResponse

	// The logging configuration for your cluster.
	Logging *Logging

	// The name of the cluster.
	Name *string

	// The platform version of your Amazon EKS cluster. For more information, see
	// Platform Versions
	// (https://docs.aws.amazon.com/eks/latest/userguide/platform-versions.html) in the
	// Amazon EKS User Guide .
	PlatformVersion *string

	// The VPC configuration used by the cluster control plane. Amazon EKS VPC
	// resources have specific requirements to work properly with Kubernetes. For more
	// information, see Cluster VPC Considerations
	// (https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html) and Cluster
	// Security Group Considerations
	// (https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the
	// Amazon EKS User Guide.
	ResourcesVpcConfig *VpcConfigResponse

	// The Amazon Resource Name (ARN) of the IAM role that provides permissions for the
	// Kubernetes control plane to make calls to Amazon Web Services API operations on
	// your behalf.
	RoleArn *string

	// The current status of the cluster.
	Status ClusterStatus

	// The metadata that you apply to the cluster to assist with categorization and
	// organization. Each tag consists of a key and an optional value. You define both.
	// Cluster tags do not propagate to any other resources associated with the
	// cluster.
	Tags map[string]string

	// The Kubernetes server version for the cluster.
	Version *string

	noSmithyDocumentSerde
}

// Compatibility information.
type Compatibility struct {

	// The supported Kubernetes version of the cluster.
	ClusterVersion *string

	// The supported default version.
	DefaultVersion bool

	// The supported compute platform.
	PlatformVersions []string

	noSmithyDocumentSerde
}

// The configuration sent to a cluster for configuration.
type ConnectorConfigRequest struct {

	// The cloud provider for the target cluster to connect.
	//
	// This member is required.
	Provider ConnectorConfigProvider

	// The Amazon Resource Name (ARN) of the role that is authorized to request the
	// connector configuration.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

// The full description of your connected cluster.
type ConnectorConfigResponse struct {

	// A unique code associated with the cluster for registration purposes.
	ActivationCode *string

	// The expiration time of the connected cluster. The cluster's YAML file must be
	// applied through the native provider.
	ActivationExpiry *time.Time

	// A unique ID associated with the cluster for registration purposes.
	ActivationId *string

	// The cluster's cloud service provider.
	Provider *string

	// The Amazon Resource Name (ARN) of the role to communicate with services from the
	// connected Kubernetes cluster.
	RoleArn *string

	noSmithyDocumentSerde
}

// The encryption configuration for the cluster.
type EncryptionConfig struct {

	// Key Management Service (KMS) key. Either the ARN or the alias can be used.
	Provider *Provider

	// Specifies the resources to be encrypted. The only supported value is "secrets".
	Resources []string

	noSmithyDocumentSerde
}

// An object representing an error when an asynchronous operation fails.
type ErrorDetail struct {

	// A brief description of the error.
	//
	// * SubnetNotFound: We couldn't find one of the
	// subnets associated with the cluster.
	//
	// * SecurityGroupNotFound: We couldn't find
	// one of the security groups associated with the cluster.
	//
	// * EniLimitReached: You
	// have reached the elastic network interface limit for your account.
	//
	// *
	// IpNotAvailable: A subnet associated with the cluster doesn't have any free IP
	// addresses.
	//
	// * AccessDenied: You don't have permissions to perform the specified
	// operation.
	//
	// * OperationNotPermitted: The service role associated with the
	// cluster doesn't have the required access permissions for Amazon EKS.
	//
	// *
	// VpcIdNotFound: We couldn't find the VPC associated with the cluster.
	ErrorCode ErrorCode

	// A more complete description of the error.
	ErrorMessage *string

	// An optional field that contains the resource IDs associated with the error.
	ResourceIds []string

	noSmithyDocumentSerde
}

// An object representing an Fargate profile.
type FargateProfile struct {

	// The name of the Amazon EKS cluster that the Fargate profile belongs to.
	ClusterName *string

	// The Unix epoch timestamp in seconds for when the Fargate profile was created.
	CreatedAt *time.Time

	// The full Amazon Resource Name (ARN) of the Fargate profile.
	FargateProfileArn *string

	// The name of the Fargate profile.
	FargateProfileName *string

	// The Amazon Resource Name (ARN) of the pod execution role to use for pods that
	// match the selectors in the Fargate profile. For more information, see Pod
	// Execution Role
	// (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in
	// the Amazon EKS User Guide.
	PodExecutionRoleArn *string

	// The selectors to match for pods to use this Fargate profile.
	Selectors []FargateProfileSelector

	// The current status of the Fargate profile.
	Status FargateProfileStatus

	// The IDs of subnets to launch pods into.
	Subnets []string

	// The metadata applied to the Fargate profile to assist with categorization and
	// organization. Each tag consists of a key and an optional value. You define both.
	// Fargate profile tags do not propagate to any other resources associated with the
	// Fargate profile, such as the pods that are scheduled with it.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An object representing an Fargate profile selector.
type FargateProfileSelector struct {

	// The Kubernetes labels that the selector should match. A pod must contain all of
	// the labels that are specified in the selector for it to be considered a match.
	Labels map[string]string

	// The Kubernetes namespace that the selector should match.
	Namespace *string

	noSmithyDocumentSerde
}

// An object representing an identity provider.
type Identity struct {

	// An object representing the OpenID Connect (https://openid.net/connect/) identity
	// provider information.
	Oidc *OIDC

	noSmithyDocumentSerde
}

// An object representing an identity provider configuration.
type IdentityProviderConfig struct {

	// The name of the identity provider configuration.
	//
	// This member is required.
	Name *string

	// The type of the identity provider configuration.
	//
	// This member is required.
	Type *string

	noSmithyDocumentSerde
}

// The full description of your identity configuration.
type IdentityProviderConfigResponse struct {

	// An object that represents an OpenID Connect (OIDC) identity provider
	// configuration.
	Oidc *OidcIdentityProviderConfig

	noSmithyDocumentSerde
}

// An object representing an issue with an Amazon EKS resource.
type Issue struct {

	// A brief description of the error.
	//
	// * AccessDenied: Amazon EKS or one or more of
	// your managed nodes is failing to authenticate or authorize with your Kubernetes
	// cluster API server.
	//
	// * AsgInstanceLaunchFailures: Your Auto Scaling group is
	// experiencing failures while attempting to launch instances.
	//
	// *
	// AutoScalingGroupNotFound: We couldn't find the Auto Scaling group associated
	// with the managed node group. You may be able to recreate an Auto Scaling group
	// with the same settings to recover.
	//
	// * ClusterUnreachable: Amazon EKS or one or
	// more of your managed nodes is unable to to communicate with your Kubernetes
	// cluster API server. This can happen if there are network disruptions or if API
	// servers are timing out processing requests.
	//
	// * Ec2LaunchTemplateNotFound: We
	// couldn't find the Amazon EC2 launch template for your managed node group. You
	// may be able to recreate a launch template with the same settings to recover.
	//
	// *
	// Ec2LaunchTemplateVersionMismatch: The Amazon EC2 launch template version for
	// your managed node group does not match the version that Amazon EKS created. You
	// may be able to revert to the version that Amazon EKS created to recover.
	//
	// *
	// Ec2SecurityGroupDeletionFailure: We could not delete the remote access security
	// group for your managed node group. Remove any dependencies from the security
	// group.
	//
	// * Ec2SecurityGroupNotFound: We couldn't find the cluster security group
	// for the cluster. You must recreate your cluster.
	//
	// *
	// Ec2SubnetInvalidConfiguration: One or more Amazon EC2 subnets specified for a
	// node group do not automatically assign public IP addresses to instances launched
	// into it. If you want your instances to be assigned a public IP address, then you
	// need to enable the auto-assign public IP address setting for the subnet. See
	// Modifying the public IPv4 addressing attribute for your subnet
	// (https://docs.aws.amazon.com/vpc/latest/userguide/vpc-ip-addressing.html#subnet-public-ip)
	// in the Amazon VPC User Guide.
	//
	// * IamInstanceProfileNotFound: We couldn't find
	// the IAM instance profile for your managed node group. You may be able to
	// recreate an instance profile with the same settings to recover.
	//
	// *
	// IamNodeRoleNotFound: We couldn't find the IAM role for your managed node group.
	// You may be able to recreate an IAM role with the same settings to recover.
	//
	// *
	// InstanceLimitExceeded: Your Amazon Web Services account is unable to launch any
	// more instances of the specified instance type. You may be able to request an
	// Amazon EC2 instance limit increase to recover.
	//
	// * InsufficientFreeAddresses: One
	// or more of the subnets associated with your managed node group does not have
	// enough available IP addresses for new nodes.
	//
	// * InternalFailure: These errors
	// are usually caused by an Amazon EKS server-side issue.
	//
	// * NodeCreationFailure:
	// Your launched instances are unable to register with your Amazon EKS cluster.
	// Common causes of this failure are insufficient node IAM role
	// (https://docs.aws.amazon.com/eks/latest/userguide/create-node-role.html)
	// permissions or lack of outbound internet access for the nodes.
	Code NodegroupIssueCode

	// The error message associated with the issue.
	Message *string

	// The Amazon Web Services resources that are afflicted by this issue.
	ResourceIds []string

	noSmithyDocumentSerde
}

// The Kubernetes network configuration for the cluster.
type KubernetesNetworkConfigRequest struct {

	// The CIDR block to assign Kubernetes service IP addresses from. If you don't
	// specify a block, Kubernetes assigns addresses from either the 10.100.0.0/16 or
	// 172.20.0.0/16 CIDR blocks. We recommend that you specify a block that does not
	// overlap with resources in other networks that are peered or connected to your
	// VPC. The block must meet the following requirements:
	//
	// * Within one of the
	// following private IP address blocks: 10.0.0.0/8, 172.16.0.0/12, or
	// 192.168.0.0/16.
	//
	// * Doesn't overlap with any CIDR block assigned to the VPC that
	// you selected for VPC.
	//
	// * Between /24 and /12.
	//
	// You can only specify a custom
	// CIDR block when you create a cluster and can't change this value once the
	// cluster is created.
	ServiceIpv4Cidr *string

	noSmithyDocumentSerde
}

// The Kubernetes network configuration for the cluster.
type KubernetesNetworkConfigResponse struct {

	// The CIDR block that Kubernetes service IP addresses are assigned from. If you
	// didn't specify a CIDR block when you created the cluster, then Kubernetes
	// assigns addresses from either the 10.100.0.0/16 or 172.20.0.0/16 CIDR blocks. If
	// this was specified, then it was specified when the cluster was created and it
	// cannot be changed.
	ServiceIpv4Cidr *string

	noSmithyDocumentSerde
}

// An object representing a node group launch template specification. The launch
// template cannot include SubnetId
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html),
// IamInstanceProfile
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IamInstanceProfile.html),
// RequestSpotInstances
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RequestSpotInstances.html),
// HibernationOptions
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_HibernationOptionsRequest.html),
// or TerminateInstances
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TerminateInstances.html),
// or the node group deployment or update will fail. For more information about
// launch templates, see CreateLaunchTemplate
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateLaunchTemplate.html)
// in the Amazon EC2 API Reference. For more information about using launch
// templates with Amazon EKS, see Launch template support
// (https://docs.aws.amazon.com/eks/latest/userguide/launch-templates.html) in the
// Amazon EKS User Guide. Specify either name or id, but not both.
type LaunchTemplateSpecification struct {

	// The ID of the launch template.
	Id *string

	// The name of the launch template.
	Name *string

	// The version of the launch template to use. If no version is specified, then the
	// template's default version is used.
	Version *string

	noSmithyDocumentSerde
}

// An object representing the logging configuration for resources in your cluster.
type Logging struct {

	// The cluster control plane logging configuration for your cluster.
	ClusterLogging []LogSetup

	noSmithyDocumentSerde
}

// An object representing the enabled or disabled Kubernetes control plane logs for
// your cluster.
type LogSetup struct {

	// If a log type is enabled, that log type exports its control plane logs to
	// CloudWatch Logs. If a log type isn't enabled, that log type doesn't export its
	// control plane logs. Each individual log type can be enabled or disabled
	// independently.
	Enabled *bool

	// The available cluster control plane log types.
	Types []LogType

	noSmithyDocumentSerde
}

// An object representing an Amazon EKS managed node group.
type Nodegroup struct {

	// If the node group was deployed using a launch template with a custom AMI, then
	// this is CUSTOM. For node groups that weren't deployed using a launch template,
	// this is the AMI type that was specified in the node group configuration.
	AmiType AMITypes

	// The capacity type of your managed node group.
	CapacityType CapacityTypes

	// The name of the cluster that the managed node group resides in.
	ClusterName *string

	// The Unix epoch timestamp in seconds for when the managed node group was created.
	CreatedAt *time.Time

	// If the node group wasn't deployed with a launch template, then this is the disk
	// size in the node group configuration. If the node group was deployed with a
	// launch template, then this is null.
	DiskSize *int32

	// The health status of the node group. If there are issues with your node group's
	// health, they are listed here.
	Health *NodegroupHealth

	// If the node group wasn't deployed with a launch template, then this is the
	// instance type that is associated with the node group. If the node group was
	// deployed with a launch template, then this is null.
	InstanceTypes []string

	// The Kubernetes labels applied to the nodes in the node group. Only labels that
	// are applied with the Amazon EKS API are shown here. There may be other
	// Kubernetes labels applied to the nodes in this group.
	Labels map[string]string

	// If a launch template was used to create the node group, then this is the launch
	// template that was used.
	LaunchTemplate *LaunchTemplateSpecification

	// The Unix epoch timestamp in seconds for when the managed node group was last
	// modified.
	ModifiedAt *time.Time

	// The IAM role associated with your node group. The Amazon EKS node kubelet daemon
	// makes calls to Amazon Web Services APIs on your behalf. Nodes receive
	// permissions for these API calls through an IAM instance profile and associated
	// policies.
	NodeRole *string

	// The Amazon Resource Name (ARN) associated with the managed node group.
	NodegroupArn *string

	// The name associated with an Amazon EKS managed node group.
	NodegroupName *string

	// If the node group was deployed using a launch template with a custom AMI, then
	// this is the AMI ID that was specified in the launch template. For node groups
	// that weren't deployed using a launch template, this is the version of the Amazon
	// EKS optimized AMI that the node group was deployed with.
	ReleaseVersion *string

	// If the node group wasn't deployed with a launch template, then this is the
	// remote access configuration that is associated with the node group. If the node
	// group was deployed with a launch template, then this is null.
	RemoteAccess *RemoteAccessConfig

	// The resources associated with the node group, such as Auto Scaling groups and
	// security groups for remote access.
	Resources *NodegroupResources

	// The scaling configuration details for the Auto Scaling group that is associated
	// with your node group.
	ScalingConfig *NodegroupScalingConfig

	// The current status of the managed node group.
	Status NodegroupStatus

	// The subnets that were specified for the Auto Scaling group that is associated
	// with your node group.
	Subnets []string

	// The metadata applied to the node group to assist with categorization and
	// organization. Each tag consists of a key and an optional value. You define both.
	// Node group tags do not propagate to any other resources associated with the node
	// group, such as the Amazon EC2 instances or subnets.
	Tags map[string]string

	// The Kubernetes taints to be applied to the nodes in the node group when they are
	// created. Effect is one of No_Schedule, Prefer_No_Schedule, or No_Execute.
	// Kubernetes taints can be used together with tolerations to control how workloads
	// are scheduled to your nodes.
	Taints []Taint

	// The node group update configuration.
	UpdateConfig *NodegroupUpdateConfig

	// The Kubernetes version of the managed node group.
	Version *string

	noSmithyDocumentSerde
}

// An object representing the health status of the node group.
type NodegroupHealth struct {

	// Any issues that are associated with the node group.
	Issues []Issue

	noSmithyDocumentSerde
}

// An object representing the resources associated with the node group, such as
// Auto Scaling groups and security groups for remote access.
type NodegroupResources struct {

	// The Auto Scaling groups associated with the node group.
	AutoScalingGroups []AutoScalingGroup

	// The remote access security group associated with the node group. This security
	// group controls SSH access to the nodes.
	RemoteAccessSecurityGroup *string

	noSmithyDocumentSerde
}

// An object representing the scaling configuration details for the Auto Scaling
// group that is associated with your node group. When creating a node group, you
// must specify all or none of the properties. When updating a node group, you can
// specify any or none of the properties.
type NodegroupScalingConfig struct {

	// The current number of nodes that the managed node group should maintain. If you
	// use Cluster Autoscaler, you shouldn't change the desiredSize value directly, as
	// this can cause the Cluster Autoscaler to suddenly scale up or scale down.
	// Whenever this parameter changes, the number of worker nodes in the node group is
	// updated to the specified size. If this parameter is given a value that is
	// smaller than the current number of running worker nodes, the necessary number of
	// worker nodes are terminated to match the given value. When using CloudFormation,
	// no action occurs if you remove this parameter from your CFN template. This
	// parameter can be different from minSize in some cases, such as when starting
	// with extra hosts for testing. This parameter can also be different when you want
	// to start with an estimated number of needed hosts, but let Cluster Autoscaler
	// reduce the number if there are too many. When Cluster Autoscaler is used, the
	// desiredSize parameter is altered by Cluster Autoscaler (but can be out-of-date
	// for short periods of time). Cluster Autoscaler doesn't scale a managed node
	// group lower than minSize or higher than maxSize.
	DesiredSize *int32

	// The maximum number of nodes that the managed node group can scale out to. For
	// information about the maximum number that you can specify, see Amazon EKS
	// service quotas
	// (https://docs.aws.amazon.com/eks/latest/userguide/service-quotas.html) in the
	// Amazon EKS User Guide.
	MaxSize *int32

	// The minimum number of nodes that the managed node group can scale in to.
	MinSize *int32

	noSmithyDocumentSerde
}

// The node group update configuration.
type NodegroupUpdateConfig struct {

	// The maximum number of nodes unavailable at once during a version update. Nodes
	// will be updated in parallel. This value or maxUnavailablePercentage is required
	// to have a value.The maximum number is 100.
	MaxUnavailable *int32

	// The maximum percentage of nodes unavailable during a version update. This
	// percentage of nodes will be updated in parallel, up to 100 nodes at once. This
	// value or maxUnavailable is required to have a value.
	MaxUnavailablePercentage *int32

	noSmithyDocumentSerde
}

// An object representing the OpenID Connect (https://openid.net/connect/) (OIDC)
// identity provider information for the cluster.
type OIDC struct {

	// The issuer URL for the OIDC identity provider.
	Issuer *string

	noSmithyDocumentSerde
}

// An object that represents the configuration for an OpenID Connect (OIDC)
// identity provider.
type OidcIdentityProviderConfig struct {

	// This is also known as audience. The ID of the client application that makes
	// authentication requests to the OIDC identity provider.
	ClientId *string

	// The cluster that the configuration is associated to.
	ClusterName *string

	// The JSON web token (JWT) claim that the provider uses to return your groups.
	GroupsClaim *string

	// The prefix that is prepended to group claims to prevent clashes with existing
	// names (such as system: groups). For example, the value oidc: creates group names
	// like oidc:engineering and oidc:infra. The prefix can't contain system:
	GroupsPrefix *string

	// The ARN of the configuration.
	IdentityProviderConfigArn *string

	// The name of the configuration.
	IdentityProviderConfigName *string

	// The URL of the OIDC identity provider that allows the API server to discover
	// public signing keys for verifying tokens.
	IssuerUrl *string

	// The key-value pairs that describe required claims in the identity token. If set,
	// each claim is verified to be present in the token with a matching value.
	RequiredClaims map[string]string

	// The status of the OIDC identity provider.
	Status ConfigStatus

	// The metadata to apply to the provider configuration to assist with
	// categorization and organization. Each tag consists of a key and an optional
	// value. You define both.
	Tags map[string]string

	// The JSON Web token (JWT) claim that is used as the username.
	UsernameClaim *string

	// The prefix that is prepended to username claims to prevent clashes with existing
	// names. The prefix can't contain system:
	UsernamePrefix *string

	noSmithyDocumentSerde
}

// An object representing an OpenID Connect (OIDC) configuration. Before
// associating an OIDC identity provider to your cluster, review the considerations
// in Authenticating users for your cluster from an OpenID Connect identity
// provider
// (https://docs.aws.amazon.com/eks/latest/userguide/authenticate-oidc-identity-provider.html)
// in the Amazon EKS User Guide.
type OidcIdentityProviderConfigRequest struct {

	// This is also known as audience. The ID for the client application that makes
	// authentication requests to the OpenID identity provider.
	//
	// This member is required.
	ClientId *string

	// The name of the OIDC provider configuration.
	//
	// This member is required.
	IdentityProviderConfigName *string

	// The URL of the OpenID identity provider that allows the API server to discover
	// public signing keys for verifying tokens. The URL must begin with https:// and
	// should correspond to the iss claim in the provider's OIDC ID tokens. Per the
	// OIDC standard, path components are allowed but query parameters are not.
	// Typically the URL consists of only a hostname, like https://server.example.org
	// or https://example.com. This URL should point to the level below
	// .well-known/openid-configuration and must be publicly accessible over the
	// internet.
	//
	// This member is required.
	IssuerUrl *string

	// The JWT claim that the provider uses to return your groups.
	GroupsClaim *string

	// The prefix that is prepended to group claims to prevent clashes with existing
	// names (such as system: groups). For example, the value oidc: will create group
	// names like oidc:engineering and oidc:infra.
	GroupsPrefix *string

	// The key value pairs that describe required claims in the identity token. If set,
	// each claim is verified to be present in the token with a matching value. For the
	// maximum number of claims that you can require, see Amazon EKS service quotas
	// (https://docs.aws.amazon.com/eks/latest/userguide/service-quotas.html) in the
	// Amazon EKS User Guide.
	RequiredClaims map[string]string

	// The JSON Web Token (JWT) claim to use as the username. The default is sub, which
	// is expected to be a unique identifier of the end user. You can choose other
	// claims, such as email or name, depending on the OpenID identity provider. Claims
	// other than email are prefixed with the issuer URL to prevent naming clashes with
	// other plug-ins.
	UsernameClaim *string

	// The prefix that is prepended to username claims to prevent clashes with existing
	// names. If you do not provide this field, and username is a value other than
	// email, the prefix defaults to issuerurl#. You can use the value - to disable all
	// prefixing.
	UsernamePrefix *string

	noSmithyDocumentSerde
}

// Identifies the Key Management Service (KMS) key used to encrypt the secrets.
type Provider struct {

	// Amazon Resource Name (ARN) or alias of the KMS key. The KMS key must be
	// symmetric, created in the same region as the cluster, and if the KMS key was
	// created in a different account, the user must have access to the KMS key. For
	// more information, see Allowing Users in Other Accounts to Use a KMS key
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-modifying-external-accounts.html)
	// in the Key Management Service Developer Guide.
	KeyArn *string

	noSmithyDocumentSerde
}

// An object representing the remote access configuration for the managed node
// group.
type RemoteAccessConfig struct {

	// The Amazon EC2 SSH key that provides access for SSH communication with the nodes
	// in the managed node group. For more information, see Amazon EC2 key pairs and
	// Linux instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html) in the
	// Amazon Elastic Compute Cloud User Guide for Linux Instances.
	Ec2SshKey *string

	// The security groups that are allowed SSH access (port 22) to the nodes. If you
	// specify an Amazon EC2 SSH key but do not specify a source security group when
	// you create a managed node group, then port 22 on the nodes is opened to the
	// internet (0.0.0.0/0). For more information, see Security Groups for Your VPC
	// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in
	// the Amazon Virtual Private Cloud User Guide.
	SourceSecurityGroups []string

	noSmithyDocumentSerde
}

// A property that allows a node to repel a set of pods.
type Taint struct {

	// The effect of the taint.
	Effect TaintEffect

	// The key of the taint.
	Key *string

	// The value of the taint.
	Value *string

	noSmithyDocumentSerde
}

// An object representing an asynchronous update.
type Update struct {

	// The Unix epoch timestamp in seconds for when the update was created.
	CreatedAt *time.Time

	// Any errors associated with a Failed update.
	Errors []ErrorDetail

	// A UUID that is used to track the update.
	Id *string

	// A key-value map that contains the parameters associated with the update.
	Params []UpdateParam

	// The current status of the update.
	Status UpdateStatus

	// The type of the update.
	Type UpdateType

	noSmithyDocumentSerde
}

// An object representing a Kubernetes label change for a managed node group.
type UpdateLabelsPayload struct {

	// Kubernetes labels to be added or updated.
	AddOrUpdateLabels map[string]string

	// Kubernetes labels to be removed.
	RemoveLabels []string

	noSmithyDocumentSerde
}

// An object representing the details of an update request.
type UpdateParam struct {

	// The keys associated with an update request.
	Type UpdateParamType

	// The value of the keys submitted as part of an update request.
	Value *string

	noSmithyDocumentSerde
}

// An object representing the details of an update to a taints payload.
type UpdateTaintsPayload struct {

	// Kubernetes taints to be added or updated.
	AddOrUpdateTaints []Taint

	// Kubernetes taints to be removed.
	RemoveTaints []Taint

	noSmithyDocumentSerde
}

// An object representing the VPC configuration to use for an Amazon EKS cluster.
type VpcConfigRequest struct {

	// Set this value to true to enable private access for your cluster's Kubernetes
	// API server endpoint. If you enable private access, Kubernetes API requests from
	// within your cluster's VPC use the private VPC endpoint. The default value for
	// this parameter is false, which disables private access for your Kubernetes API
	// server. If you disable private access and you have nodes or Fargate pods in the
	// cluster, then ensure that publicAccessCidrs includes the necessary CIDR blocks
	// for communication with the nodes or Fargate pods. For more information, see
	// Amazon EKS cluster endpoint access control
	// (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the
	// Amazon EKS User Guide .
	EndpointPrivateAccess *bool

	// Set this value to false to disable public access to your cluster's Kubernetes
	// API server endpoint. If you disable public access, your cluster's Kubernetes API
	// server can only receive requests from within the cluster VPC. The default value
	// for this parameter is true, which enables public access for your Kubernetes API
	// server. For more information, see Amazon EKS cluster endpoint access control
	// (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the
	// Amazon EKS User Guide .
	EndpointPublicAccess *bool

	// The CIDR blocks that are allowed access to your cluster's public Kubernetes API
	// server endpoint. Communication to the endpoint from addresses outside of the
	// CIDR blocks that you specify is denied. The default value is 0.0.0.0/0. If
	// you've disabled private endpoint access and you have nodes or Fargate pods in
	// the cluster, then ensure that you specify the necessary CIDR blocks. For more
	// information, see Amazon EKS cluster endpoint access control
	// (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the
	// Amazon EKS User Guide .
	PublicAccessCidrs []string

	// Specify one or more security groups for the cross-account elastic network
	// interfaces that Amazon EKS creates to use that allow communication between your
	// nodes and the Kubernetes control plane. If you don't specify any security
	// groups, then familiarize yourself with the difference between Amazon EKS
	// defaults for clusters deployed with Kubernetes:
	//
	// * 1.14 Amazon EKS platform
	// version eks.2 and earlier
	//
	// * 1.14 Amazon EKS platform version eks.3 and
	// later
	//
	// For more information, see Amazon EKS security group considerations
	// (https://docs.aws.amazon.com/eks/latest/userguide/sec-group-reqs.html) in the
	// Amazon EKS User Guide .
	SecurityGroupIds []string

	// Specify subnets for your Amazon EKS nodes. Amazon EKS creates cross-account
	// elastic network interfaces in these subnets to allow communication between your
	// nodes and the Kubernetes control plane.
	SubnetIds []string

	noSmithyDocumentSerde
}

// An object representing an Amazon EKS cluster VPC configuration response.
type VpcConfigResponse struct {

	// The cluster security group that was created by Amazon EKS for the cluster.
	// Managed node groups use this security group for control-plane-to-data-plane
	// communication.
	ClusterSecurityGroupId *string

	// This parameter indicates whether the Amazon EKS private API server endpoint is
	// enabled. If the Amazon EKS private API server endpoint is enabled, Kubernetes
	// API requests that originate from within your cluster's VPC use the private VPC
	// endpoint instead of traversing the internet. If this value is disabled and you
	// have nodes or Fargate pods in the cluster, then ensure that publicAccessCidrs
	// includes the necessary CIDR blocks for communication with the nodes or Fargate
	// pods. For more information, see Amazon EKS cluster endpoint access control
	// (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the
	// Amazon EKS User Guide .
	EndpointPrivateAccess bool

	// This parameter indicates whether the Amazon EKS public API server endpoint is
	// enabled. If the Amazon EKS public API server endpoint is disabled, your
	// cluster's Kubernetes API server can only receive requests that originate from
	// within the cluster VPC.
	EndpointPublicAccess bool

	// The CIDR blocks that are allowed access to your cluster's public Kubernetes API
	// server endpoint. Communication to the endpoint from addresses outside of the
	// listed CIDR blocks is denied. The default value is 0.0.0.0/0. If you've disabled
	// private endpoint access and you have nodes or Fargate pods in the cluster, then
	// ensure that the necessary CIDR blocks are listed. For more information, see
	// Amazon EKS cluster endpoint access control
	// (https://docs.aws.amazon.com/eks/latest/userguide/cluster-endpoint.html) in the
	// Amazon EKS User Guide .
	PublicAccessCidrs []string

	// The security groups associated with the cross-account elastic network interfaces
	// that are used to allow communication between your nodes and the Kubernetes
	// control plane.
	SecurityGroupIds []string

	// The subnets associated with your cluster.
	SubnetIds []string

	// The VPC associated with your cluster.
	VpcId *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
