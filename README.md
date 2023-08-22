# go-sdk-v2-example
This is go-sdk-v2 sample

## Usage
You should check go version and install sdk for go v2(It is for 1.15 or later)
```bash
go mod init main

go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/ec2
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
  DryRun *bool => true or false (default false),
  # There are many FilterTypes (cidr, cidr-block-association.cidr-block, state, owner-id, is-default, vpc-id, ...etc)
  Filters []types.Filter,
  # The maximum number of items to return
  MaxResults *int32,
  # The token returned from a previous paginated req
  NextToken *string,
  # One or more VPC ids (default all VPCs)
  VpcIds []string
}
```
`DescribeVpcsOutput`
```bash
{
  NextToken *string,
  # Information about one or more vpc
  Vpcs []types.Vpc,
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
