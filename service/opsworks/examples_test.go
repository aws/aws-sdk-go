package opsworks_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/opsworks"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleOpsWorks_AssignInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.AssignInstanceInput{
		InstanceID: aws.String("String"), // Required
		LayerIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.AssignInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_AssignVolume() {
	svc := opsworks.New(nil)

	params := &opsworks.AssignVolumeInput{
		VolumeID:   aws.String("String"), // Required
		InstanceID: aws.String("String"),
	}
	resp, err := svc.AssignVolume(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_AssociateElasticIP() {
	svc := opsworks.New(nil)

	params := &opsworks.AssociateElasticIPInput{
		ElasticIP:  aws.String("String"), // Required
		InstanceID: aws.String("String"),
	}
	resp, err := svc.AssociateElasticIP(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_AttachElasticLoadBalancer() {
	svc := opsworks.New(nil)

	params := &opsworks.AttachElasticLoadBalancerInput{
		ElasticLoadBalancerName: aws.String("String"), // Required
		LayerID:                 aws.String("String"), // Required
	}
	resp, err := svc.AttachElasticLoadBalancer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_CloneStack() {
	svc := opsworks.New(nil)

	params := &opsworks.CloneStackInput{
		ServiceRoleARN: aws.String("String"), // Required
		SourceStackID:  aws.String("String"), // Required
		Attributes: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		ChefConfiguration: &opsworks.ChefConfiguration{
			BerkshelfVersion: aws.String("String"),
			ManageBerkshelf:  aws.Boolean(true),
		},
		CloneAppIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		ClonePermissions: aws.Boolean(true),
		ConfigurationManager: &opsworks.StackConfigurationManager{
			Name:    aws.String("String"),
			Version: aws.String("String"),
		},
		CustomCookbooksSource: &opsworks.Source{
			Password: aws.String("String"),
			Revision: aws.String("String"),
			SSHKey:   aws.String("String"),
			Type:     aws.String("SourceType"),
			URL:      aws.String("String"),
			Username: aws.String("String"),
		},
		CustomJSON:                aws.String("String"),
		DefaultAvailabilityZone:   aws.String("String"),
		DefaultInstanceProfileARN: aws.String("String"),
		DefaultOs:                 aws.String("String"),
		DefaultRootDeviceType:     aws.String("RootDeviceType"),
		DefaultSSHKeyName:         aws.String("String"),
		DefaultSubnetID:           aws.String("String"),
		HostnameTheme:             aws.String("String"),
		Name:                      aws.String("String"),
		Region:                    aws.String("String"),
		UseCustomCookbooks:        aws.Boolean(true),
		UseOpsWorksSecurityGroups: aws.Boolean(true),
		VPCID: aws.String("String"),
	}
	resp, err := svc.CloneStack(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_CreateApp() {
	svc := opsworks.New(nil)

	params := &opsworks.CreateAppInput{
		Name:    aws.String("String"),  // Required
		StackID: aws.String("String"),  // Required
		Type:    aws.String("AppType"), // Required
		AppSource: &opsworks.Source{
			Password: aws.String("String"),
			Revision: aws.String("String"),
			SSHKey:   aws.String("String"),
			Type:     aws.String("SourceType"),
			URL:      aws.String("String"),
			Username: aws.String("String"),
		},
		Attributes: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		DataSources: []*opsworks.DataSource{
			&opsworks.DataSource{ // Required
				ARN:          aws.String("String"),
				DatabaseName: aws.String("String"),
				Type:         aws.String("String"),
			},
			// More values...
		},
		Description: aws.String("String"),
		Domains: []*string{
			aws.String("String"), // Required
			// More values...
		},
		EnableSSL: aws.Boolean(true),
		Environment: []*opsworks.EnvironmentVariable{
			&opsworks.EnvironmentVariable{ // Required
				Key:    aws.String("String"), // Required
				Value:  aws.String("String"), // Required
				Secure: aws.Boolean(true),
			},
			// More values...
		},
		SSLConfiguration: &opsworks.SSLConfiguration{
			Certificate: aws.String("String"), // Required
			PrivateKey:  aws.String("String"), // Required
			Chain:       aws.String("String"),
		},
		Shortname: aws.String("String"),
	}
	resp, err := svc.CreateApp(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_CreateDeployment() {
	svc := opsworks.New(nil)

	params := &opsworks.CreateDeploymentInput{
		Command: &opsworks.DeploymentCommand{ // Required
			Name: aws.String("DeploymentCommandName"), // Required
			Args: &map[string][]*string{
				"Key": []*string{ // Required
					aws.String("String"), // Required
					// More values...
				},
				// More values...
			},
		},
		StackID:    aws.String("String"), // Required
		AppID:      aws.String("String"),
		Comment:    aws.String("String"),
		CustomJSON: aws.String("String"),
		InstanceIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.CreateDeployment(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_CreateInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.CreateInstanceInput{
		InstanceType: aws.String("String"), // Required
		LayerIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
		StackID:              aws.String("String"), // Required
		AMIID:                aws.String("String"),
		Architecture:         aws.String("Architecture"),
		AutoScalingType:      aws.String("AutoScalingType"),
		AvailabilityZone:     aws.String("String"),
		EBSOptimized:         aws.Boolean(true),
		Hostname:             aws.String("String"),
		InstallUpdatesOnBoot: aws.Boolean(true),
		Os:                   aws.String("String"),
		RootDeviceType:       aws.String("RootDeviceType"),
		SSHKeyName:           aws.String("String"),
		SubnetID:             aws.String("String"),
		VirtualizationType:   aws.String("String"),
	}
	resp, err := svc.CreateInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_CreateLayer() {
	svc := opsworks.New(nil)

	params := &opsworks.CreateLayerInput{
		Name:      aws.String("String"),    // Required
		Shortname: aws.String("String"),    // Required
		StackID:   aws.String("String"),    // Required
		Type:      aws.String("LayerType"), // Required
		Attributes: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		AutoAssignElasticIPs:     aws.Boolean(true),
		AutoAssignPublicIPs:      aws.Boolean(true),
		CustomInstanceProfileARN: aws.String("String"),
		CustomRecipes: &opsworks.Recipes{
			Configure: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Deploy: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Setup: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Shutdown: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Undeploy: []*string{
				aws.String("String"), // Required
				// More values...
			},
		},
		CustomSecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		EnableAutoHealing:    aws.Boolean(true),
		InstallUpdatesOnBoot: aws.Boolean(true),
		LifecycleEventConfiguration: &opsworks.LifecycleEventConfiguration{
			Shutdown: &opsworks.ShutdownEventConfiguration{
				DelayUntilELBConnectionsDrained: aws.Boolean(true),
				ExecutionTimeout:                aws.Long(1),
			},
		},
		Packages: []*string{
			aws.String("String"), // Required
			// More values...
		},
		UseEBSOptimizedInstances: aws.Boolean(true),
		VolumeConfigurations: []*opsworks.VolumeConfiguration{
			&opsworks.VolumeConfiguration{ // Required
				MountPoint:    aws.String("String"), // Required
				NumberOfDisks: aws.Long(1),          // Required
				Size:          aws.Long(1),          // Required
				IOPS:          aws.Long(1),
				RAIDLevel:     aws.Long(1),
				VolumeType:    aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.CreateLayer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_CreateStack() {
	svc := opsworks.New(nil)

	params := &opsworks.CreateStackInput{
		DefaultInstanceProfileARN: aws.String("String"), // Required
		Name:           aws.String("String"), // Required
		Region:         aws.String("String"), // Required
		ServiceRoleARN: aws.String("String"), // Required
		Attributes: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		ChefConfiguration: &opsworks.ChefConfiguration{
			BerkshelfVersion: aws.String("String"),
			ManageBerkshelf:  aws.Boolean(true),
		},
		ConfigurationManager: &opsworks.StackConfigurationManager{
			Name:    aws.String("String"),
			Version: aws.String("String"),
		},
		CustomCookbooksSource: &opsworks.Source{
			Password: aws.String("String"),
			Revision: aws.String("String"),
			SSHKey:   aws.String("String"),
			Type:     aws.String("SourceType"),
			URL:      aws.String("String"),
			Username: aws.String("String"),
		},
		CustomJSON:                aws.String("String"),
		DefaultAvailabilityZone:   aws.String("String"),
		DefaultOs:                 aws.String("String"),
		DefaultRootDeviceType:     aws.String("RootDeviceType"),
		DefaultSSHKeyName:         aws.String("String"),
		DefaultSubnetID:           aws.String("String"),
		HostnameTheme:             aws.String("String"),
		UseCustomCookbooks:        aws.Boolean(true),
		UseOpsWorksSecurityGroups: aws.Boolean(true),
		VPCID: aws.String("String"),
	}
	resp, err := svc.CreateStack(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_CreateUserProfile() {
	svc := opsworks.New(nil)

	params := &opsworks.CreateUserProfileInput{
		IAMUserARN:          aws.String("String"), // Required
		AllowSelfManagement: aws.Boolean(true),
		SSHPublicKey:        aws.String("String"),
		SSHUsername:         aws.String("String"),
	}
	resp, err := svc.CreateUserProfile(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DeleteApp() {
	svc := opsworks.New(nil)

	params := &opsworks.DeleteAppInput{
		AppID: aws.String("String"), // Required
	}
	resp, err := svc.DeleteApp(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DeleteInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.DeleteInstanceInput{
		InstanceID:      aws.String("String"), // Required
		DeleteElasticIP: aws.Boolean(true),
		DeleteVolumes:   aws.Boolean(true),
	}
	resp, err := svc.DeleteInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DeleteLayer() {
	svc := opsworks.New(nil)

	params := &opsworks.DeleteLayerInput{
		LayerID: aws.String("String"), // Required
	}
	resp, err := svc.DeleteLayer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DeleteStack() {
	svc := opsworks.New(nil)

	params := &opsworks.DeleteStackInput{
		StackID: aws.String("String"), // Required
	}
	resp, err := svc.DeleteStack(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DeleteUserProfile() {
	svc := opsworks.New(nil)

	params := &opsworks.DeleteUserProfileInput{
		IAMUserARN: aws.String("String"), // Required
	}
	resp, err := svc.DeleteUserProfile(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DeregisterElasticIP() {
	svc := opsworks.New(nil)

	params := &opsworks.DeregisterElasticIPInput{
		ElasticIP: aws.String("String"), // Required
	}
	resp, err := svc.DeregisterElasticIP(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DeregisterInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.DeregisterInstanceInput{
		InstanceID: aws.String("String"), // Required
	}
	resp, err := svc.DeregisterInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DeregisterRDSDBInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.DeregisterRDSDBInstanceInput{
		RDSDBInstanceARN: aws.String("String"), // Required
	}
	resp, err := svc.DeregisterRDSDBInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DeregisterVolume() {
	svc := opsworks.New(nil)

	params := &opsworks.DeregisterVolumeInput{
		VolumeID: aws.String("String"), // Required
	}
	resp, err := svc.DeregisterVolume(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeApps() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeAppsInput{
		AppIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		StackID: aws.String("String"),
	}
	resp, err := svc.DescribeApps(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeCommands() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeCommandsInput{
		CommandIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		DeploymentID: aws.String("String"),
		InstanceID:   aws.String("String"),
	}
	resp, err := svc.DescribeCommands(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeDeployments() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeDeploymentsInput{
		AppID: aws.String("String"),
		DeploymentIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		StackID: aws.String("String"),
	}
	resp, err := svc.DescribeDeployments(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeElasticIPs() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeElasticIPsInput{
		IPs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		InstanceID: aws.String("String"),
		StackID:    aws.String("String"),
	}
	resp, err := svc.DescribeElasticIPs(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeElasticLoadBalancers() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeElasticLoadBalancersInput{
		LayerIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		StackID: aws.String("String"),
	}
	resp, err := svc.DescribeElasticLoadBalancers(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeInstances() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeInstancesInput{
		InstanceIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		LayerID: aws.String("String"),
		StackID: aws.String("String"),
	}
	resp, err := svc.DescribeInstances(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeLayers() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeLayersInput{
		LayerIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		StackID: aws.String("String"),
	}
	resp, err := svc.DescribeLayers(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeLoadBasedAutoScaling() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeLoadBasedAutoScalingInput{
		LayerIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeLoadBasedAutoScaling(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeMyUserProfile() {
	svc := opsworks.New(nil)

	var params *opsworks.DescribeMyUserProfileInput
	resp, err := svc.DescribeMyUserProfile(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribePermissions() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribePermissionsInput{
		IAMUserARN: aws.String("String"),
		StackID:    aws.String("String"),
	}
	resp, err := svc.DescribePermissions(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeRAIDArrays() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeRAIDArraysInput{
		InstanceID: aws.String("String"),
		RAIDArrayIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		StackID: aws.String("String"),
	}
	resp, err := svc.DescribeRAIDArrays(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeRDSDBInstances() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeRDSDBInstancesInput{
		StackID: aws.String("String"), // Required
		RDSDBInstanceARNs: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeRDSDBInstances(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeServiceErrors() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeServiceErrorsInput{
		InstanceID: aws.String("String"),
		ServiceErrorIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		StackID: aws.String("String"),
	}
	resp, err := svc.DescribeServiceErrors(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeStackProvisioningParameters() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeStackProvisioningParametersInput{
		StackID: aws.String("String"), // Required
	}
	resp, err := svc.DescribeStackProvisioningParameters(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeStackSummary() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeStackSummaryInput{
		StackID: aws.String("String"), // Required
	}
	resp, err := svc.DescribeStackSummary(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeStacks() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeStacksInput{
		StackIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeStacks(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeTimeBasedAutoScaling() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeTimeBasedAutoScalingInput{
		InstanceIDs: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeTimeBasedAutoScaling(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeUserProfiles() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeUserProfilesInput{
		IAMUserARNs: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeUserProfiles(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DescribeVolumes() {
	svc := opsworks.New(nil)

	params := &opsworks.DescribeVolumesInput{
		InstanceID:  aws.String("String"),
		RAIDArrayID: aws.String("String"),
		StackID:     aws.String("String"),
		VolumeIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeVolumes(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DetachElasticLoadBalancer() {
	svc := opsworks.New(nil)

	params := &opsworks.DetachElasticLoadBalancerInput{
		ElasticLoadBalancerName: aws.String("String"), // Required
		LayerID:                 aws.String("String"), // Required
	}
	resp, err := svc.DetachElasticLoadBalancer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_DisassociateElasticIP() {
	svc := opsworks.New(nil)

	params := &opsworks.DisassociateElasticIPInput{
		ElasticIP: aws.String("String"), // Required
	}
	resp, err := svc.DisassociateElasticIP(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_GetHostnameSuggestion() {
	svc := opsworks.New(nil)

	params := &opsworks.GetHostnameSuggestionInput{
		LayerID: aws.String("String"), // Required
	}
	resp, err := svc.GetHostnameSuggestion(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_RebootInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.RebootInstanceInput{
		InstanceID: aws.String("String"), // Required
	}
	resp, err := svc.RebootInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_RegisterElasticIP() {
	svc := opsworks.New(nil)

	params := &opsworks.RegisterElasticIPInput{
		ElasticIP: aws.String("String"), // Required
		StackID:   aws.String("String"), // Required
	}
	resp, err := svc.RegisterElasticIP(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_RegisterInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.RegisterInstanceInput{
		StackID:  aws.String("String"), // Required
		Hostname: aws.String("String"),
		InstanceIdentity: &opsworks.InstanceIdentity{
			Document:  aws.String("String"),
			Signature: aws.String("String"),
		},
		PrivateIP:               aws.String("String"),
		PublicIP:                aws.String("String"),
		RSAPublicKey:            aws.String("String"),
		RSAPublicKeyFingerprint: aws.String("String"),
	}
	resp, err := svc.RegisterInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_RegisterRDSDBInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.RegisterRDSDBInstanceInput{
		DBPassword:       aws.String("String"), // Required
		DBUser:           aws.String("String"), // Required
		RDSDBInstanceARN: aws.String("String"), // Required
		StackID:          aws.String("String"), // Required
	}
	resp, err := svc.RegisterRDSDBInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_RegisterVolume() {
	svc := opsworks.New(nil)

	params := &opsworks.RegisterVolumeInput{
		StackID:     aws.String("String"), // Required
		EC2VolumeID: aws.String("String"),
	}
	resp, err := svc.RegisterVolume(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_SetLoadBasedAutoScaling() {
	svc := opsworks.New(nil)

	params := &opsworks.SetLoadBasedAutoScalingInput{
		LayerID: aws.String("String"), // Required
		DownScaling: &opsworks.AutoScalingThresholds{
			CPUThreshold:       aws.Double(1.0),
			IgnoreMetricsTime:  aws.Long(1),
			InstanceCount:      aws.Long(1),
			LoadThreshold:      aws.Double(1.0),
			MemoryThreshold:    aws.Double(1.0),
			ThresholdsWaitTime: aws.Long(1),
		},
		Enable: aws.Boolean(true),
		UpScaling: &opsworks.AutoScalingThresholds{
			CPUThreshold:       aws.Double(1.0),
			IgnoreMetricsTime:  aws.Long(1),
			InstanceCount:      aws.Long(1),
			LoadThreshold:      aws.Double(1.0),
			MemoryThreshold:    aws.Double(1.0),
			ThresholdsWaitTime: aws.Long(1),
		},
	}
	resp, err := svc.SetLoadBasedAutoScaling(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_SetPermission() {
	svc := opsworks.New(nil)

	params := &opsworks.SetPermissionInput{
		IAMUserARN: aws.String("String"), // Required
		StackID:    aws.String("String"), // Required
		AllowSSH:   aws.Boolean(true),
		AllowSudo:  aws.Boolean(true),
		Level:      aws.String("String"),
	}
	resp, err := svc.SetPermission(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_SetTimeBasedAutoScaling() {
	svc := opsworks.New(nil)

	params := &opsworks.SetTimeBasedAutoScalingInput{
		InstanceID: aws.String("String"), // Required
		AutoScalingSchedule: &opsworks.WeeklyAutoScalingSchedule{
			Friday: &map[string]*string{
				"Key": aws.String("Switch"), // Required
				// More values...
			},
			Monday: &map[string]*string{
				"Key": aws.String("Switch"), // Required
				// More values...
			},
			Saturday: &map[string]*string{
				"Key": aws.String("Switch"), // Required
				// More values...
			},
			Sunday: &map[string]*string{
				"Key": aws.String("Switch"), // Required
				// More values...
			},
			Thursday: &map[string]*string{
				"Key": aws.String("Switch"), // Required
				// More values...
			},
			Tuesday: &map[string]*string{
				"Key": aws.String("Switch"), // Required
				// More values...
			},
			Wednesday: &map[string]*string{
				"Key": aws.String("Switch"), // Required
				// More values...
			},
		},
	}
	resp, err := svc.SetTimeBasedAutoScaling(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_StartInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.StartInstanceInput{
		InstanceID: aws.String("String"), // Required
	}
	resp, err := svc.StartInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_StartStack() {
	svc := opsworks.New(nil)

	params := &opsworks.StartStackInput{
		StackID: aws.String("String"), // Required
	}
	resp, err := svc.StartStack(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_StopInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.StopInstanceInput{
		InstanceID: aws.String("String"), // Required
	}
	resp, err := svc.StopInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_StopStack() {
	svc := opsworks.New(nil)

	params := &opsworks.StopStackInput{
		StackID: aws.String("String"), // Required
	}
	resp, err := svc.StopStack(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UnassignInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.UnassignInstanceInput{
		InstanceID: aws.String("String"), // Required
	}
	resp, err := svc.UnassignInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UnassignVolume() {
	svc := opsworks.New(nil)

	params := &opsworks.UnassignVolumeInput{
		VolumeID: aws.String("String"), // Required
	}
	resp, err := svc.UnassignVolume(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UpdateApp() {
	svc := opsworks.New(nil)

	params := &opsworks.UpdateAppInput{
		AppID: aws.String("String"), // Required
		AppSource: &opsworks.Source{
			Password: aws.String("String"),
			Revision: aws.String("String"),
			SSHKey:   aws.String("String"),
			Type:     aws.String("SourceType"),
			URL:      aws.String("String"),
			Username: aws.String("String"),
		},
		Attributes: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		DataSources: []*opsworks.DataSource{
			&opsworks.DataSource{ // Required
				ARN:          aws.String("String"),
				DatabaseName: aws.String("String"),
				Type:         aws.String("String"),
			},
			// More values...
		},
		Description: aws.String("String"),
		Domains: []*string{
			aws.String("String"), // Required
			// More values...
		},
		EnableSSL: aws.Boolean(true),
		Environment: []*opsworks.EnvironmentVariable{
			&opsworks.EnvironmentVariable{ // Required
				Key:    aws.String("String"), // Required
				Value:  aws.String("String"), // Required
				Secure: aws.Boolean(true),
			},
			// More values...
		},
		Name: aws.String("String"),
		SSLConfiguration: &opsworks.SSLConfiguration{
			Certificate: aws.String("String"), // Required
			PrivateKey:  aws.String("String"), // Required
			Chain:       aws.String("String"),
		},
		Type: aws.String("AppType"),
	}
	resp, err := svc.UpdateApp(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UpdateElasticIP() {
	svc := opsworks.New(nil)

	params := &opsworks.UpdateElasticIPInput{
		ElasticIP: aws.String("String"), // Required
		Name:      aws.String("String"),
	}
	resp, err := svc.UpdateElasticIP(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UpdateInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.UpdateInstanceInput{
		InstanceID:           aws.String("String"), // Required
		AMIID:                aws.String("String"),
		Architecture:         aws.String("Architecture"),
		AutoScalingType:      aws.String("AutoScalingType"),
		EBSOptimized:         aws.Boolean(true),
		Hostname:             aws.String("String"),
		InstallUpdatesOnBoot: aws.Boolean(true),
		InstanceType:         aws.String("String"),
		LayerIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		Os:         aws.String("String"),
		SSHKeyName: aws.String("String"),
	}
	resp, err := svc.UpdateInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UpdateLayer() {
	svc := opsworks.New(nil)

	params := &opsworks.UpdateLayerInput{
		LayerID: aws.String("String"), // Required
		Attributes: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		AutoAssignElasticIPs:     aws.Boolean(true),
		AutoAssignPublicIPs:      aws.Boolean(true),
		CustomInstanceProfileARN: aws.String("String"),
		CustomRecipes: &opsworks.Recipes{
			Configure: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Deploy: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Setup: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Shutdown: []*string{
				aws.String("String"), // Required
				// More values...
			},
			Undeploy: []*string{
				aws.String("String"), // Required
				// More values...
			},
		},
		CustomSecurityGroupIDs: []*string{
			aws.String("String"), // Required
			// More values...
		},
		EnableAutoHealing:    aws.Boolean(true),
		InstallUpdatesOnBoot: aws.Boolean(true),
		LifecycleEventConfiguration: &opsworks.LifecycleEventConfiguration{
			Shutdown: &opsworks.ShutdownEventConfiguration{
				DelayUntilELBConnectionsDrained: aws.Boolean(true),
				ExecutionTimeout:                aws.Long(1),
			},
		},
		Name: aws.String("String"),
		Packages: []*string{
			aws.String("String"), // Required
			// More values...
		},
		Shortname:                aws.String("String"),
		UseEBSOptimizedInstances: aws.Boolean(true),
		VolumeConfigurations: []*opsworks.VolumeConfiguration{
			&opsworks.VolumeConfiguration{ // Required
				MountPoint:    aws.String("String"), // Required
				NumberOfDisks: aws.Long(1),          // Required
				Size:          aws.Long(1),          // Required
				IOPS:          aws.Long(1),
				RAIDLevel:     aws.Long(1),
				VolumeType:    aws.String("String"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateLayer(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UpdateMyUserProfile() {
	svc := opsworks.New(nil)

	params := &opsworks.UpdateMyUserProfileInput{
		SSHPublicKey: aws.String("String"),
	}
	resp, err := svc.UpdateMyUserProfile(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UpdateRDSDBInstance() {
	svc := opsworks.New(nil)

	params := &opsworks.UpdateRDSDBInstanceInput{
		RDSDBInstanceARN: aws.String("String"), // Required
		DBPassword:       aws.String("String"),
		DBUser:           aws.String("String"),
	}
	resp, err := svc.UpdateRDSDBInstance(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UpdateStack() {
	svc := opsworks.New(nil)

	params := &opsworks.UpdateStackInput{
		StackID: aws.String("String"), // Required
		Attributes: &map[string]*string{
			"Key": aws.String("String"), // Required
			// More values...
		},
		ChefConfiguration: &opsworks.ChefConfiguration{
			BerkshelfVersion: aws.String("String"),
			ManageBerkshelf:  aws.Boolean(true),
		},
		ConfigurationManager: &opsworks.StackConfigurationManager{
			Name:    aws.String("String"),
			Version: aws.String("String"),
		},
		CustomCookbooksSource: &opsworks.Source{
			Password: aws.String("String"),
			Revision: aws.String("String"),
			SSHKey:   aws.String("String"),
			Type:     aws.String("SourceType"),
			URL:      aws.String("String"),
			Username: aws.String("String"),
		},
		CustomJSON:                aws.String("String"),
		DefaultAvailabilityZone:   aws.String("String"),
		DefaultInstanceProfileARN: aws.String("String"),
		DefaultOs:                 aws.String("String"),
		DefaultRootDeviceType:     aws.String("RootDeviceType"),
		DefaultSSHKeyName:         aws.String("String"),
		DefaultSubnetID:           aws.String("String"),
		HostnameTheme:             aws.String("String"),
		Name:                      aws.String("String"),
		ServiceRoleARN:            aws.String("String"),
		UseCustomCookbooks:        aws.Boolean(true),
		UseOpsWorksSecurityGroups: aws.Boolean(true),
	}
	resp, err := svc.UpdateStack(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UpdateUserProfile() {
	svc := opsworks.New(nil)

	params := &opsworks.UpdateUserProfileInput{
		IAMUserARN:          aws.String("String"), // Required
		AllowSelfManagement: aws.Boolean(true),
		SSHPublicKey:        aws.String("String"),
		SSHUsername:         aws.String("String"),
	}
	resp, err := svc.UpdateUserProfile(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleOpsWorks_UpdateVolume() {
	svc := opsworks.New(nil)

	params := &opsworks.UpdateVolumeInput{
		VolumeID:   aws.String("String"), // Required
		MountPoint: aws.String("String"),
		Name:       aws.String("String"),
	}
	resp, err := svc.UpdateVolume(params)

	if awsErr, ok := err.(awserr.Error); ok {
		// Generic AWS Error with Code, Message, and original error (if any)
		fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			// A service error occurred
			fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
		}
	} else {
		fmt.Println(err.Error())
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}