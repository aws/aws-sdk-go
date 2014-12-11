package aws

import (
	"net/url"
	"reflect"
	"testing"
)

type Filter struct {
	Name   StringValue `ec2:"Name" xml:"Name"`
	Values []string    `ec2:"Value" xml:"Value>item"`
}

type DescribeImagesRequest struct {
	DryRun          BooleanValue `ec2:"dryRun" xml:"dryRun"`
	ExecutableUsers []string     `ec2:"ExecutableBy" xml:"ExecutableBy>ExecutableBy"`
	Filters         []Filter     `ec2:"Filter" xml:"Filter>Filter"`
	ImageIds        []string     `ec2:"ImageId" xml:"ImageId>ImageId"`
	Owners          []string     `ec2:"Owner" xml:"Owner>Owner"`
}

func TestSerializingEC2QueryRequests(t *testing.T) {
	c := EC2Client{}
	actual := url.Values{}
	req := DescribeImagesRequest{
		Filters: []Filter{
			{Name: String("owner-id"), Values: []string{"yay"}},
			{Name: String("is-public"), Values: []string{"true"}},
		},
	}
	if err := c.loadValues(actual, req, ""); err != nil {
		t.Fatal(err)
	}

	expected := url.Values{
		"Filter.1.Name":    []string{"owner-id"},
		"Filter.1.Value.1": []string{"yay"},
		"Filter.2.Name":    []string{"is-public"},
		"Filter.2.Value.1": []string{"true"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Was \n%s\n but expected \n%s", actual.Encode(), expected.Encode())
	}
}

type AutoScalingGroupNamesType struct {
	AutoScalingGroupNames []string     `xml:"AutoScalingGroupNames>member"`
	MaxRecords            IntegerValue `xml:"MaxRecords"`
	NextToken             StringValue  `xml:"NextToken"`
}

func TestSerializingNonEC2QueryRequests(t *testing.T) {
	c := QueryClient{}
	actual := url.Values{}
	req := AutoScalingGroupNamesType{
		AutoScalingGroupNames: []string{"one", "two"},
		MaxRecords:            Integer(100),
		NextToken:             String("wobble"),
	}
	if err := c.loadValues(actual, req, ""); err != nil {
		t.Fatal(err)
	}

	expected := url.Values{
		"AutoScalingGroupNames.member.1": []string{"one"},
		"AutoScalingGroupNames.member.2": []string{"two"},
		"MaxRecords":                     []string{"100"},
		"NextToken":                      []string{"wobble"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Was %#v \n but expected \n %#v", actual, expected)
	}
}
