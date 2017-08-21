
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

