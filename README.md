# go-sdk-v2-example
This is go-sdk-v2 sample
https://github.com/aws/aws-sdk-go-v2
## Usage
You should check go version and install sdk for go v2(It is for 1.15 or later)
```bash
go mod init main

go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/ec2
go get github.com/aws/aws-sdk-go-v2/service/eks
go get github.com/gofiber/fiber/v2
```
Also, you should get your aws keys(access,secret). There are many ways to store keys. I stored `~/.aws/credentials` like this.
```bash
[default]
aws_access_key_id = your_access_key
aws_secret_access_key = your_secret_key
```
`context.TODO()` will find your access key in your credentials directory.

## GetVpc
### Request
`GET /vpcs`
```bash
    curl -i -H 'Accept: application/json' http://localhost:3000/vpcs
```
### Response
vpcs data of DescribeVpcsOutput
```JSON
[
  {
    "CidrBlock":"172.31.0.0/16",
    "CidrBlockAssociationSet":
    [
      {
        "AssociationId":"vpcAssociationId",
        "CidrBlock":"172.31.0.0/16",
        "CidrBlockState":{"State":"associated","StatusMessage":null}
      }
    ],
    "DhcpOptionsId":"dopt-0baf9efe196339cea",
    "InstanceTenancy":"default",
    "Ipv6CidrBlockAssociationSet":null,
    "IsDefault":true,
    "OwnerId":"yourAWSId",
    "State":"available",
    "Tags":null,
    "VpcId":"vpcId"
  }
]
```

`DescribeVpcssInput`

```bash
{
  # Check whether you have the required permissions for the action
  DryRun *bool # => true or false (default false)

  # There are many FilterTypes (cidr, cidr-block-association.cidr-block, state, owner-id, is-default, vpc-id, ...etc)
  Filters []types.Filter

  # The maximum number of items to return
  MaxResults *int32

  # The token returned from a previous paginated req
  NextToken *string

  # One or more VPC ids (default all VPCs)
  VpcIds []string
}
```
`DescribeVpcsOutput`
```bash
{
  NextToken *string

  # Information about one or more vpc
  Vpcs []types.Vpc

  # Metadata about operation's result (req ID, ...etc)
  ResultMetadata middleware.Metadata
}
```


## GetVPCSubnets

### Request
`GET /vpcSubnets`
```bash
    curl -i -H 'Accept: application/json' http://localhost:3000/vpcSubnets
```
### Response
Subnets data of DescribeSubnetsOutput
```JSON
[
  {
    "AssignIpv6AddressOnCreation":false,
    "AvailabilityZone":"us-west-2a",
    "AvailabilityZoneId":"usw2-az1",
    "AvailableIpAddressCount":4091,
    "CidrBlock":"172.31.16.0/20",
    "CustomerOwnedIpv4Pool":null,
    "DefaultForAz":true,
    "EnableDns64":false,
    "EnableLniAtDeviceIndex":null,
    "Ipv6CidrBlockAssociationSet":[],
    "Ipv6Native":false,
    "MapCustomerOwnedIpOnLaunch":false,
    "MapPublicIpOnLaunch":true,
    "OutpostArn":null,
    "OwnerId":"yourAWSId",
    "PrivateDnsNameOptionsOnLaunch":{"EnableResourceNameDnsAAAARecord":false,"EnableResourceNameDnsARecord":false,"HostnameType":"ip-name"},
    "State":"available",
    "SubnetArn":"subnetArn",
    "SubnetId":"subnetId",
    "Tags":null,
    "VpcId":"vpcId"
  }
]
```

`DescribeSubnetsInput`
```bash
{
  #Check whether you have the required permissions for the action
  DryRun *bool => true or false (default false),

  # There are many FilterTypes (availability-zone, availability-zone-id, available-ip-address-count,cidr-block,subnet-arn,subnet-id,state, ...etc)
  Filters []types.Filter,

  # The maximum number of items to return
  MaxResults *int32,

  # The token returned from a previous paginated req
  NextToken *string,

  # One or more subnet ids (default all subnets)
  SubnetIds []string
}
```
`DescribeSubnetsOutput`
```bash
{
  NextToken *string,

  # Information about one or more subnets
  Subnets []types.Subnet,

  # Metadata about operation's result (req ID, ...etc)
  ResultMetadata middleware.Metadata
}
```

## GetEC2InstanceTypes

### Request
`GET /instanceTypes`
```bash
    curl -i -H 'Accept: application/json' http://localhost:3000/instanceTypes
```
### Response
InstanceTypes data of DescribeInstanceTypesOutput
```JSON
[
  {
    "AutoRecoverySupported":true,
    "BareMetal":false,
    "BurstablePerformanceSupported":false,
    "CurrentGeneration":false,
    "DedicatedHostsSupported":false,
    "EbsInfo":
    {
      "EbsOptimizedInfo":null,
      "EbsOptimizedSupport":"unsupported",
      "EncryptionSupport":"supported",
      "NvmeSupport":"unsupported"
    },
    "FpgaInfo":null,
    "FreeTierEligible":true,
    "GpuInfo":null,
    "HibernationSupported":false,
    "Hypervisor":"xen",
    "InferenceAcceleratorInfo":null,
    "InstanceStorageInfo":null,
    "InstanceStorageSupported":false,
    "InstanceType":"t1.micro",
    "MemoryInfo":{"SizeInMiB":627},
    "NetworkInfo":
    {
      "DefaultNetworkCardIndex":0,
      "EfaInfo":null,
      "EfaSupported":false,
      "EnaSrdSupported":false,
      "EnaSupport":"unsupported",
      "EncryptionInTransitSupported":false,
      "Ipv4AddressesPerInterface":2,
      "Ipv6AddressesPerInterface":0,
      "Ipv6Supported":false,
      "MaximumNetworkCards":1,
      "MaximumNetworkInterfaces":2,
      "NetworkCards":[{"MaximumNetworkInterfaces":2,"NetworkCardIndex":0,"NetworkPerformance":"Very Low"}],
      "NetworkPerformance":"Very Low"
    },
    "PlacementGroupInfo":{"SupportedStrategies":["partition","spread"]},
    "ProcessorInfo":{"SupportedArchitectures":["i386","x86_64"],"SupportedFeatures":null,"SustainedClockSpeedInGhz":null},
    "SupportedBootModes":["legacy-bios"],
    "SupportedRootDeviceTypes":["ebs"],
    "SupportedUsageClasses":["on-demand","spot"],
    "SupportedVirtualizationTypes":["hvm","paravirtual"],
    "VCpuInfo":{"DefaultCores":1,"DefaultThreadsPerCore":1,"DefaultVCpus":1,"ValidCores":null,"ValidThreadsPerCore":null}
  }
]
```

`DescribeInstanceTypesInput`
```bash
{
  # Check whether you have the required permissions for the action
  DryRun *bool => true or false (default false),
  # There are many FilterTypes (bare-metal, free-tier-eligible, instance-type(using '*'), ebs-info.*, instance-storage-info.*, network-info.* ...etc)
  Filters []types.Filter,
  # The maximum number of items to return
  MaxResults *int32,
  # The token returned from a previous paginated req
  NextToken *string,
  # One or more instance types (default all instance types)
  InstanceTypes []types.InstanceType
}
```
`DescribeInstanceTypesOutput`
```bash
{
  NextToken *string,
  # The instance Type
  InstanceTypes []types.InstanceTypeInfo,
  # Metadata about operation's result (req ID, ...etc)
  ResultMetadata middleware.Metadata
}
```

## GetEC2AMI

### Request
`GET /ami`
```bash
    curl -i -H 'Accept: application/json' http://localhost:3000/ami
```
### Response
Image data of DescribeImagesOutput
```JSON
[
  {
    "Architecture":"x86_64",
    "BlockDeviceMappings":
    [
      {
        "DeviceName":"/dev/sda1",
        "Ebs":
        {
          "DeleteOnTermination":true,
          "Encrypted":false,
          "Iops":null,
          "KmsKeyId":null,
          "OutpostArn":null,
          "SnapshotId":"snapshotId",
          "Throughput":null,
          "VolumeSize":50,
          "VolumeType":"gp2"
        },
        "NoDevice":null,
        "VirtualName":null
      },
    ],
    "BootMode":"",
    "CreationDate":"ImageCreationDate",
    "DeprecationTime":"ImageDeprecationDate",
    "Description":"ImageDescription",
    "EnaSupport":true,
    "Hypervisor":"xen",
    "ImageId":"amiID",
    "ImageLocation":"ImageLocation",
    "ImageOwnerAlias":"amazon",
    "ImageType":"machine",
    "ImdsSupport":"",
    "KernelId":null,
    "Name":"ImageName",
    "OwnerId":"OwnerId",
    "Platform":"windows",
    "PlatformDetails":"Windows",
    "ProductCodes":null,
    "Public":true,
    "RamdiskId":null,
    "RootDeviceName":"/dev/sda1",
    "RootDeviceType":"ebs",
    "SriovNetSupport":"simple",
    "State":"available",
    "StateReason":null,
    "Tags":null,
    "TpmSupport":"",
    "UsageOperation":"RunInstances:0002",
    "VirtualizationType":"hvm"
  }
]
```

`DescribeImagesInput`
```bash
{
  # Check whether you have the required permissions for the action
  DryRun *bool # => true or false (default false)

  # Specify an AWS account ID, and then AMIs shared with taht specific AWS account ID are returned.
  ExecutableUsers []string

  # There are many FilterTypes (architecture, block-device-mapping.*, creation-date, hypervisor, image-id, product-code ...etc). If you want to see details, please check the link(https://github.com/aws/aws-sdk-go-v2/blob/main/service/ec2/api_op_DescribeImages.go)
  Filters []types.Filter

  # The image IDs, default all images
  ImageIds []string

  # Whether to include deprecated AMIs, default is not included.
  IncludeDeprecated *bool

  # The maximum number of items to return
  MaxResults *int32

  # The token returned from a previous paginated req
  NextToken *string

  # Specify specified images owners, you enter AWS account IDs, self, amazon, aws-marketplace
  Owners []string
}
```
`DescribeImagesOutput`
```bash
{
  NextToken *string,

  # Information about the images
  Images []types.Image

  # Metadata about operation's result (req ID, ...etc)
  ResultMetadata middleware.Metadata
}
```

## GetEKSClusterVersion
It is temporary way, so If you want to find supported cluster version, find another way.
### Request
`GET /eksClusterVersion`
```bash
    curl -i -H 'Accept: application/json' http://localhost:3000/eksClusterVersion
```
### Response
Cluster version list of EKS (using eksaddon info)
```JSON
["1.27","1.26","1.25","1.24","1.23","1.22","1.21","1.20"]
```

`DescribeAddonVersionsInput`
```bash
{
  # The name of add-on. Add-on extend the functionality of k8s.(https://kubernetes.io/docs/concepts/cluster-administration/addons/)
  AddonName *string

  # The k8s version that you can use the add-on with. The example uses this value.
  KubernetesVersion *string

  # The maximum number of items to return
  MaxResults *int32,

  # The token returned from a previous paginated req
  NextToken *string,

  # Specify specified images owners, you enter AWS account IDs, self, amazon, aws-marketplace
  Owners []string

  # The publisher of add-on
  Publishers []string

  # The type of the add-on
	Types []string
}
```
`DescribeAddonVersionsOutput`
```bash
{
  NextToken *string,

  # The list of available versions with k8s version compatibility
  Addons []types.AddonInfo

  # Metadata about operation's result (req ID, ...etc)
  ResultMetadata middleware.Metadata
}
```

## GetCostUsage

### Request
`GET /costUsage`
```bash
    curl -i -H 'Accept: application/json' http://localhost:3000/costUsage
```
### Response
The time period that's covered by the results (Used Resource)
ResultsByTime data of GetCostAndUsageOutput
```JSON
[
  {
    "Estimated": false,
    "Groups": [
      {
        "Keys": [
          "AWS Data Transfer"
        ],
        "Metrics": {
          "BlendedCost": {
            "Amount": "-0.000000074",
            "Unit": "USD"
          }
        }
      },
      {
        "Keys": [
          "AWS Key Management Service"
        ],
        "Metrics": {
          "BlendedCost": {
            "Amount": "0",
            "Unit": "USD"
          }
        }
      }, ...
    ],
    "TimePeriod": {
      "End": "2023-10-01",
      "Start": "2023-09-30"
    },
    "Total": {}
  }
]
```

`GetCostAndUsageInput`
```bash
{
  # You can set the date interval. It is required and you can choose either MONTHLY or DAILY or HOURLY.
  Granularity types.Granularity

  # It is about cost type like AmortizedCost, BlendedCost, NetAmortizedCost, NetUnblendedCost, NormalizedUsageAmount, UnblendedCost, and UsageQuantity.
  # It is required and if you want to see costs of all accounts, then choose BlendedCost.
  Metrics []string

  # You can set the time interval from start to end. It it required.
  TimePeriod *types.DateInterval

  # You can filter costs by different dimensions. (FYI; https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
  Filter *types.Expression

  # You can group costs by different dimensions, or tag keys or cost categories.
  GroupBy []types.GroupDefinition

  # It is for paging to retrieve next results.
 NextPageToken *string

}
```
`GetCostAndUsageOutput`
```bash
{
  # The attributes that apply to a specific dimension value. Usually, it is empty if you don't set it.
  DimensionValueAttributes []types.DimensionValuesWithAttributes

	# It is about your group filter.
	GroupDefinitions []types.GroupDefinition

	# It is the token for the next set of retrievable results.
	NextPageToken *string

	# The time period that's covered by the results in the response. You can make your cost chart using this.
	ResultsByTime []types.ResultByTime

	# Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
```
