# Example

This is an example using the AWS SDK for Go to list ec2 instances that match provided tag name filter.


# Usage

The example uses the bucket name provided, and lists all object keys in a bucket.

```sh
go run -tags example filter_ec2_by_tag.go <name_filter>
```

Output:
```
listing instances with tag vpn in: us-east-1
[{
  Instances: [{
      AmiLaunchIndex: 0,
      Architecture: "x86_64",
      BlockDeviceMappings: [{
          DeviceName: "/dev/xvda",
          Ebs: {
            AttachTime: 2016-07-06 18:04:53 +0000 UTC,
            DeleteOnTermination: true,
            Status: "attached",
            VolumeId: "vol-xxxx"
          }
        }],
      ...
```

# Example  Fetch By region

This is an example using the AWS SDK for Go to list ec2 instances instance state By different region . By default it fetch all running and stopped instance 


# Usage


```sh
go run  filter_ec2_by_region running
```

Output:
```

Fetching instace details  for region:  ap-south-1
printing instance details.....
instance id i-00cf3fcsssdd373766
current State stopped
done for region ap-south-1 ****



Fetching instace details  for region:  eu-west-2
There is no instance for the for region:  eu-west-2
done for region eu-west-2 ****

