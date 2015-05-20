package swf_test

import (
	"bytes"
	"fmt"
	"time"
	"github.com/awslabs/aws-sdk-go/aws"

	"github.com/awslabs/aws-sdk-go/aws/awserr"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/swf"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleSWF_CountClosedWorkflowExecutions() {
	svc := swf.New(nil)

	params := &swf.CountClosedWorkflowExecutionsInput{
		Domain: aws.String("DomainName"), // Required
		CloseStatusFilter: &swf.CloseStatusFilter{
			Status: aws.String("CloseStatus"), // Required
		},
		CloseTimeFilter: &swf.ExecutionTimeFilter{
			OldestDate: aws.Time(time.Now()), // Required
			LatestDate: aws.Time(time.Now()),
		},
		ExecutionFilter: &swf.WorkflowExecutionFilter{
			WorkflowID: aws.String("WorkflowId"), // Required
		},
		StartTimeFilter: &swf.ExecutionTimeFilter{
			OldestDate: aws.Time(time.Now()), // Required
			LatestDate: aws.Time(time.Now()),
		},
		TagFilter: &swf.TagFilter{
			Tag: aws.String("Tag"), // Required
		},
		TypeFilter: &swf.WorkflowTypeFilter{
			Name:    aws.String("Name"), // Required
			Version: aws.String("VersionOptional"),
		},
	}
	resp, err := svc.CountClosedWorkflowExecutions(params)

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

func ExampleSWF_CountOpenWorkflowExecutions() {
	svc := swf.New(nil)

	params := &swf.CountOpenWorkflowExecutionsInput{
		Domain: aws.String("DomainName"), // Required
		StartTimeFilter: &swf.ExecutionTimeFilter{ // Required
			OldestDate: aws.Time(time.Now()), // Required
			LatestDate: aws.Time(time.Now()),
		},
		ExecutionFilter: &swf.WorkflowExecutionFilter{
			WorkflowID: aws.String("WorkflowId"), // Required
		},
		TagFilter: &swf.TagFilter{
			Tag: aws.String("Tag"), // Required
		},
		TypeFilter: &swf.WorkflowTypeFilter{
			Name:    aws.String("Name"), // Required
			Version: aws.String("VersionOptional"),
		},
	}
	resp, err := svc.CountOpenWorkflowExecutions(params)

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

func ExampleSWF_CountPendingActivityTasks() {
	svc := swf.New(nil)

	params := &swf.CountPendingActivityTasksInput{
		Domain: aws.String("DomainName"), // Required
		TaskList: &swf.TaskList{ // Required
			Name: aws.String("Name"), // Required
		},
	}
	resp, err := svc.CountPendingActivityTasks(params)

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

func ExampleSWF_CountPendingDecisionTasks() {
	svc := swf.New(nil)

	params := &swf.CountPendingDecisionTasksInput{
		Domain: aws.String("DomainName"), // Required
		TaskList: &swf.TaskList{ // Required
			Name: aws.String("Name"), // Required
		},
	}
	resp, err := svc.CountPendingDecisionTasks(params)

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

func ExampleSWF_DeprecateActivityType() {
	svc := swf.New(nil)

	params := &swf.DeprecateActivityTypeInput{
		ActivityType: &swf.ActivityType{ // Required
			Name:    aws.String("Name"),    // Required
			Version: aws.String("Version"), // Required
		},
		Domain: aws.String("DomainName"), // Required
	}
	resp, err := svc.DeprecateActivityType(params)

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

func ExampleSWF_DeprecateDomain() {
	svc := swf.New(nil)

	params := &swf.DeprecateDomainInput{
		Name: aws.String("DomainName"), // Required
	}
	resp, err := svc.DeprecateDomain(params)

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

func ExampleSWF_DeprecateWorkflowType() {
	svc := swf.New(nil)

	params := &swf.DeprecateWorkflowTypeInput{
		Domain: aws.String("DomainName"), // Required
		WorkflowType: &swf.WorkflowType{ // Required
			Name:    aws.String("Name"),    // Required
			Version: aws.String("Version"), // Required
		},
	}
	resp, err := svc.DeprecateWorkflowType(params)

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

func ExampleSWF_DescribeActivityType() {
	svc := swf.New(nil)

	params := &swf.DescribeActivityTypeInput{
		ActivityType: &swf.ActivityType{ // Required
			Name:    aws.String("Name"),    // Required
			Version: aws.String("Version"), // Required
		},
		Domain: aws.String("DomainName"), // Required
	}
	resp, err := svc.DescribeActivityType(params)

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

func ExampleSWF_DescribeDomain() {
	svc := swf.New(nil)

	params := &swf.DescribeDomainInput{
		Name: aws.String("DomainName"), // Required
	}
	resp, err := svc.DescribeDomain(params)

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

func ExampleSWF_DescribeWorkflowExecution() {
	svc := swf.New(nil)

	params := &swf.DescribeWorkflowExecutionInput{
		Domain: aws.String("DomainName"), // Required
		Execution: &swf.WorkflowExecution{ // Required
			RunID:      aws.String("RunId"),      // Required
			WorkflowID: aws.String("WorkflowId"), // Required
		},
	}
	resp, err := svc.DescribeWorkflowExecution(params)

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

func ExampleSWF_DescribeWorkflowType() {
	svc := swf.New(nil)

	params := &swf.DescribeWorkflowTypeInput{
		Domain: aws.String("DomainName"), // Required
		WorkflowType: &swf.WorkflowType{ // Required
			Name:    aws.String("Name"),    // Required
			Version: aws.String("Version"), // Required
		},
	}
	resp, err := svc.DescribeWorkflowType(params)

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

func ExampleSWF_GetWorkflowExecutionHistory() {
	svc := swf.New(nil)

	params := &swf.GetWorkflowExecutionHistoryInput{
		Domain: aws.String("DomainName"), // Required
		Execution: &swf.WorkflowExecution{ // Required
			RunID:      aws.String("RunId"),      // Required
			WorkflowID: aws.String("WorkflowId"), // Required
		},
		MaximumPageSize: aws.Long(1),
		NextPageToken:   aws.String("PageToken"),
		ReverseOrder:    aws.Boolean(true),
	}
	resp, err := svc.GetWorkflowExecutionHistory(params)

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

func ExampleSWF_ListActivityTypes() {
	svc := swf.New(nil)

	params := &swf.ListActivityTypesInput{
		Domain:             aws.String("DomainName"),         // Required
		RegistrationStatus: aws.String("RegistrationStatus"), // Required
		MaximumPageSize:    aws.Long(1),
		Name:               aws.String("Name"),
		NextPageToken:      aws.String("PageToken"),
		ReverseOrder:       aws.Boolean(true),
	}
	resp, err := svc.ListActivityTypes(params)

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

func ExampleSWF_ListClosedWorkflowExecutions() {
	svc := swf.New(nil)

	params := &swf.ListClosedWorkflowExecutionsInput{
		Domain: aws.String("DomainName"), // Required
		CloseStatusFilter: &swf.CloseStatusFilter{
			Status: aws.String("CloseStatus"), // Required
		},
		CloseTimeFilter: &swf.ExecutionTimeFilter{
			OldestDate: aws.Time(time.Now()), // Required
			LatestDate: aws.Time(time.Now()),
		},
		ExecutionFilter: &swf.WorkflowExecutionFilter{
			WorkflowID: aws.String("WorkflowId"), // Required
		},
		MaximumPageSize: aws.Long(1),
		NextPageToken:   aws.String("PageToken"),
		ReverseOrder:    aws.Boolean(true),
		StartTimeFilter: &swf.ExecutionTimeFilter{
			OldestDate: aws.Time(time.Now()), // Required
			LatestDate: aws.Time(time.Now()),
		},
		TagFilter: &swf.TagFilter{
			Tag: aws.String("Tag"), // Required
		},
		TypeFilter: &swf.WorkflowTypeFilter{
			Name:    aws.String("Name"), // Required
			Version: aws.String("VersionOptional"),
		},
	}
	resp, err := svc.ListClosedWorkflowExecutions(params)

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

func ExampleSWF_ListDomains() {
	svc := swf.New(nil)

	params := &swf.ListDomainsInput{
		RegistrationStatus: aws.String("RegistrationStatus"), // Required
		MaximumPageSize:    aws.Long(1),
		NextPageToken:      aws.String("PageToken"),
		ReverseOrder:       aws.Boolean(true),
	}
	resp, err := svc.ListDomains(params)

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

func ExampleSWF_ListOpenWorkflowExecutions() {
	svc := swf.New(nil)

	params := &swf.ListOpenWorkflowExecutionsInput{
		Domain: aws.String("DomainName"), // Required
		StartTimeFilter: &swf.ExecutionTimeFilter{ // Required
			OldestDate: aws.Time(time.Now()), // Required
			LatestDate: aws.Time(time.Now()),
		},
		ExecutionFilter: &swf.WorkflowExecutionFilter{
			WorkflowID: aws.String("WorkflowId"), // Required
		},
		MaximumPageSize: aws.Long(1),
		NextPageToken:   aws.String("PageToken"),
		ReverseOrder:    aws.Boolean(true),
		TagFilter: &swf.TagFilter{
			Tag: aws.String("Tag"), // Required
		},
		TypeFilter: &swf.WorkflowTypeFilter{
			Name:    aws.String("Name"), // Required
			Version: aws.String("VersionOptional"),
		},
	}
	resp, err := svc.ListOpenWorkflowExecutions(params)

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

func ExampleSWF_ListWorkflowTypes() {
	svc := swf.New(nil)

	params := &swf.ListWorkflowTypesInput{
		Domain:             aws.String("DomainName"),         // Required
		RegistrationStatus: aws.String("RegistrationStatus"), // Required
		MaximumPageSize:    aws.Long(1),
		Name:               aws.String("Name"),
		NextPageToken:      aws.String("PageToken"),
		ReverseOrder:       aws.Boolean(true),
	}
	resp, err := svc.ListWorkflowTypes(params)

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

func ExampleSWF_PollForActivityTask() {
	svc := swf.New(nil)

	params := &swf.PollForActivityTaskInput{
		Domain: aws.String("DomainName"), // Required
		TaskList: &swf.TaskList{ // Required
			Name: aws.String("Name"), // Required
		},
		Identity: aws.String("Identity"),
	}
	resp, err := svc.PollForActivityTask(params)

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

func ExampleSWF_PollForDecisionTask() {
	svc := swf.New(nil)

	params := &swf.PollForDecisionTaskInput{
		Domain: aws.String("DomainName"), // Required
		TaskList: &swf.TaskList{ // Required
			Name: aws.String("Name"), // Required
		},
		Identity:        aws.String("Identity"),
		MaximumPageSize: aws.Long(1),
		NextPageToken:   aws.String("PageToken"),
		ReverseOrder:    aws.Boolean(true),
	}
	resp, err := svc.PollForDecisionTask(params)

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

func ExampleSWF_RecordActivityTaskHeartbeat() {
	svc := swf.New(nil)

	params := &swf.RecordActivityTaskHeartbeatInput{
		TaskToken: aws.String("TaskToken"), // Required
		Details:   aws.String("LimitedData"),
	}
	resp, err := svc.RecordActivityTaskHeartbeat(params)

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

func ExampleSWF_RegisterActivityType() {
	svc := swf.New(nil)

	params := &swf.RegisterActivityTypeInput{
		Domain:                      aws.String("DomainName"), // Required
		Name:                        aws.String("Name"),       // Required
		Version:                     aws.String("Version"),    // Required
		DefaultTaskHeartbeatTimeout: aws.String("DurationInSecondsOptional"),
		DefaultTaskList: &swf.TaskList{
			Name: aws.String("Name"), // Required
		},
		DefaultTaskPriority:               aws.String("TaskPriority"),
		DefaultTaskScheduleToCloseTimeout: aws.String("DurationInSecondsOptional"),
		DefaultTaskScheduleToStartTimeout: aws.String("DurationInSecondsOptional"),
		DefaultTaskStartToCloseTimeout:    aws.String("DurationInSecondsOptional"),
		Description:                       aws.String("Description"),
	}
	resp, err := svc.RegisterActivityType(params)

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

func ExampleSWF_RegisterDomain() {
	svc := swf.New(nil)

	params := &swf.RegisterDomainInput{
		Name: aws.String("DomainName"), // Required
		WorkflowExecutionRetentionPeriodInDays: aws.String("DurationInDays"), // Required
		Description:                            aws.String("Description"),
	}
	resp, err := svc.RegisterDomain(params)

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

func ExampleSWF_RegisterWorkflowType() {
	svc := swf.New(nil)

	params := &swf.RegisterWorkflowTypeInput{
		Domain:                              aws.String("DomainName"), // Required
		Name:                                aws.String("Name"),       // Required
		Version:                             aws.String("Version"),    // Required
		DefaultChildPolicy:                  aws.String("ChildPolicy"),
		DefaultExecutionStartToCloseTimeout: aws.String("DurationInSecondsOptional"),
		DefaultTaskList: &swf.TaskList{
			Name: aws.String("Name"), // Required
		},
		DefaultTaskPriority:            aws.String("TaskPriority"),
		DefaultTaskStartToCloseTimeout: aws.String("DurationInSecondsOptional"),
		Description:                    aws.String("Description"),
	}
	resp, err := svc.RegisterWorkflowType(params)

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

func ExampleSWF_RequestCancelWorkflowExecution() {
	svc := swf.New(nil)

	params := &swf.RequestCancelWorkflowExecutionInput{
		Domain:     aws.String("DomainName"), // Required
		WorkflowID: aws.String("WorkflowId"), // Required
		RunID:      aws.String("RunIdOptional"),
	}
	resp, err := svc.RequestCancelWorkflowExecution(params)

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

func ExampleSWF_RespondActivityTaskCanceled() {
	svc := swf.New(nil)

	params := &swf.RespondActivityTaskCanceledInput{
		TaskToken: aws.String("TaskToken"), // Required
		Details:   aws.String("Data"),
	}
	resp, err := svc.RespondActivityTaskCanceled(params)

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

func ExampleSWF_RespondActivityTaskCompleted() {
	svc := swf.New(nil)

	params := &swf.RespondActivityTaskCompletedInput{
		TaskToken: aws.String("TaskToken"), // Required
		Result:    aws.String("Data"),
	}
	resp, err := svc.RespondActivityTaskCompleted(params)

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

func ExampleSWF_RespondActivityTaskFailed() {
	svc := swf.New(nil)

	params := &swf.RespondActivityTaskFailedInput{
		TaskToken: aws.String("TaskToken"), // Required
		Details:   aws.String("Data"),
		Reason:    aws.String("FailureReason"),
	}
	resp, err := svc.RespondActivityTaskFailed(params)

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

func ExampleSWF_RespondDecisionTaskCompleted() {
	svc := swf.New(nil)

	params := &swf.RespondDecisionTaskCompletedInput{
		TaskToken: aws.String("TaskToken"), // Required
		Decisions: []*swf.Decision{
			&swf.Decision{ // Required
				DecisionType: aws.String("DecisionType"), // Required
				CancelTimerDecisionAttributes: &swf.CancelTimerDecisionAttributes{
					TimerID: aws.String("TimerId"), // Required
				},
				CancelWorkflowExecutionDecisionAttributes: &swf.CancelWorkflowExecutionDecisionAttributes{
					Details: aws.String("Data"),
				},
				CompleteWorkflowExecutionDecisionAttributes: &swf.CompleteWorkflowExecutionDecisionAttributes{
					Result: aws.String("Data"),
				},
				ContinueAsNewWorkflowExecutionDecisionAttributes: &swf.ContinueAsNewWorkflowExecutionDecisionAttributes{
					ChildPolicy:                  aws.String("ChildPolicy"),
					ExecutionStartToCloseTimeout: aws.String("DurationInSecondsOptional"),
					Input: aws.String("Data"),
					TagList: []*string{
						aws.String("Tag"), // Required
						// More values...
					},
					TaskList: &swf.TaskList{
						Name: aws.String("Name"), // Required
					},
					TaskPriority:            aws.String("TaskPriority"),
					TaskStartToCloseTimeout: aws.String("DurationInSecondsOptional"),
					WorkflowTypeVersion:     aws.String("Version"),
				},
				FailWorkflowExecutionDecisionAttributes: &swf.FailWorkflowExecutionDecisionAttributes{
					Details: aws.String("Data"),
					Reason:  aws.String("FailureReason"),
				},
				RecordMarkerDecisionAttributes: &swf.RecordMarkerDecisionAttributes{
					MarkerName: aws.String("MarkerName"), // Required
					Details:    aws.String("Data"),
				},
				RequestCancelActivityTaskDecisionAttributes: &swf.RequestCancelActivityTaskDecisionAttributes{
					ActivityID: aws.String("ActivityId"), // Required
				},
				RequestCancelExternalWorkflowExecutionDecisionAttributes: &swf.RequestCancelExternalWorkflowExecutionDecisionAttributes{
					WorkflowID: aws.String("WorkflowId"), // Required
					Control:    aws.String("Data"),
					RunID:      aws.String("RunIdOptional"),
				},
				ScheduleActivityTaskDecisionAttributes: &swf.ScheduleActivityTaskDecisionAttributes{
					ActivityID: aws.String("ActivityId"), // Required
					ActivityType: &swf.ActivityType{ // Required
						Name:    aws.String("Name"),    // Required
						Version: aws.String("Version"), // Required
					},
					Control:          aws.String("Data"),
					HeartbeatTimeout: aws.String("DurationInSecondsOptional"),
					Input:            aws.String("Data"),
					ScheduleToCloseTimeout: aws.String("DurationInSecondsOptional"),
					ScheduleToStartTimeout: aws.String("DurationInSecondsOptional"),
					StartToCloseTimeout:    aws.String("DurationInSecondsOptional"),
					TaskList: &swf.TaskList{
						Name: aws.String("Name"), // Required
					},
					TaskPriority: aws.String("TaskPriority"),
				},
				SignalExternalWorkflowExecutionDecisionAttributes: &swf.SignalExternalWorkflowExecutionDecisionAttributes{
					SignalName: aws.String("SignalName"), // Required
					WorkflowID: aws.String("WorkflowId"), // Required
					Control:    aws.String("Data"),
					Input:      aws.String("Data"),
					RunID:      aws.String("RunIdOptional"),
				},
				StartChildWorkflowExecutionDecisionAttributes: &swf.StartChildWorkflowExecutionDecisionAttributes{
					WorkflowID: aws.String("WorkflowId"), // Required
					WorkflowType: &swf.WorkflowType{ // Required
						Name:    aws.String("Name"),    // Required
						Version: aws.String("Version"), // Required
					},
					ChildPolicy: aws.String("ChildPolicy"),
					Control:     aws.String("Data"),
					ExecutionStartToCloseTimeout: aws.String("DurationInSecondsOptional"),
					Input: aws.String("Data"),
					TagList: []*string{
						aws.String("Tag"), // Required
						// More values...
					},
					TaskList: &swf.TaskList{
						Name: aws.String("Name"), // Required
					},
					TaskPriority:            aws.String("TaskPriority"),
					TaskStartToCloseTimeout: aws.String("DurationInSecondsOptional"),
				},
				StartTimerDecisionAttributes: &swf.StartTimerDecisionAttributes{
					StartToFireTimeout: aws.String("DurationInSeconds"), // Required
					TimerID:            aws.String("TimerId"),           // Required
					Control:            aws.String("Data"),
				},
			},
			// More values...
		},
		ExecutionContext: aws.String("Data"),
	}
	resp, err := svc.RespondDecisionTaskCompleted(params)

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

func ExampleSWF_SignalWorkflowExecution() {
	svc := swf.New(nil)

	params := &swf.SignalWorkflowExecutionInput{
		Domain:     aws.String("DomainName"), // Required
		SignalName: aws.String("SignalName"), // Required
		WorkflowID: aws.String("WorkflowId"), // Required
		Input:      aws.String("Data"),
		RunID:      aws.String("RunIdOptional"),
	}
	resp, err := svc.SignalWorkflowExecution(params)

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

func ExampleSWF_StartWorkflowExecution() {
	svc := swf.New(nil)

	params := &swf.StartWorkflowExecutionInput{
		Domain:     aws.String("DomainName"), // Required
		WorkflowID: aws.String("WorkflowId"), // Required
		WorkflowType: &swf.WorkflowType{ // Required
			Name:    aws.String("Name"),    // Required
			Version: aws.String("Version"), // Required
		},
		ChildPolicy:                  aws.String("ChildPolicy"),
		ExecutionStartToCloseTimeout: aws.String("DurationInSecondsOptional"),
		Input: aws.String("Data"),
		TagList: []*string{
			aws.String("Tag"), // Required
			// More values...
		},
		TaskList: &swf.TaskList{
			Name: aws.String("Name"), // Required
		},
		TaskPriority:            aws.String("TaskPriority"),
		TaskStartToCloseTimeout: aws.String("DurationInSecondsOptional"),
	}
	resp, err := svc.StartWorkflowExecution(params)

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

func ExampleSWF_TerminateWorkflowExecution() {
	svc := swf.New(nil)

	params := &swf.TerminateWorkflowExecutionInput{
		Domain:      aws.String("DomainName"), // Required
		WorkflowID:  aws.String("WorkflowId"), // Required
		ChildPolicy: aws.String("ChildPolicy"),
		Details:     aws.String("Data"),
		Reason:      aws.String("TerminateReason"),
		RunID:       aws.String("RunIdOptional"),
	}
	resp, err := svc.TerminateWorkflowExecution(params)

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