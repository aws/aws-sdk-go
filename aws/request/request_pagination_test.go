package request_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting"
)

func TestPagination_Standalone(t *testing.T) {
	type testPageInput struct {
		NextToken *string
	}
	type testPageOutput struct {
		Value     *string
		NextToken *string
	}
	type testCase struct {
		Value, PrevToken, NextToken *string
	}

	type testCaseList struct {
		StopOnSameToken bool
		Cases           []testCase
	}

	cases := []testCaseList{
		{
			Cases: []testCase{
				{aws.String("FirstValue"), aws.String("InitalToken"), aws.String("FirstToken")},
				{aws.String("SecondValue"), aws.String("FirstToken"), aws.String("SecondToken")},
				{aws.String("ThirdValue"), aws.String("SecondToken"), nil},
			},
			StopOnSameToken: false,
		},
		{
			Cases: []testCase{
				{aws.String("FirstValue"), aws.String("InitalToken"), aws.String("FirstToken")},
				{aws.String("SecondValue"), aws.String("FirstToken"), aws.String("SecondToken")},
				{aws.String("ThirdValue"), aws.String("SecondToken"), aws.String("")},
			},
			StopOnSameToken: false,
		},
		{
			Cases: []testCase{
				{aws.String("FirstValue"), aws.String("InitalToken"), aws.String("FirstToken")},
				{aws.String("SecondValue"), aws.String("FirstToken"), aws.String("SecondToken")},
				{nil, aws.String("SecondToken"), aws.String("SecondToken")},
			},
			StopOnSameToken: true,
		},
		{
			Cases: []testCase{
				{aws.String("FirstValue"), aws.String("InitalToken"), aws.String("FirstToken")},
				{aws.String("SecondValue"), aws.String("FirstToken"), aws.String("SecondToken")},
				{aws.String("SecondValue"), aws.String("SecondToken"), aws.String("SecondToken")},
			},
			StopOnSameToken: true,
		},
	}

	for _, testcase := range cases {
		c := testcase.Cases
		input := testPageInput{
			NextToken: c[0].PrevToken,
		}

		svc := awstesting.NewClient()
		i := 0
		p := request.Pagination{
			EndPageOnSameToken: testcase.StopOnSameToken,
			NewRequest: func() (*request.Request, error) {
				r := svc.NewRequest(
					&request.Operation{
						Name: "Operation",
						Paginator: &request.Paginator{
							InputTokens:  []string{"NextToken"},
							OutputTokens: []string{"NextToken"},
						},
					},
					&input, &testPageOutput{},
				)
				// Setup handlers for testing
				r.Handlers.Clear()
				r.Handlers.Build.PushBack(func(req *request.Request) {
					if e, a := len(c), i+1; a > e {
						t.Fatalf("expect no more than %d requests, got %d", e, a)
					}
					in := req.Params.(*testPageInput)
					if e, a := aws.StringValue(c[i].PrevToken), aws.StringValue(in.NextToken); e != a {
						t.Errorf("%d, expect NextToken input %q, got %q", i, e, a)
					}
				})
				r.Handlers.Unmarshal.PushBack(func(req *request.Request) {
					out := &testPageOutput{
						Value: c[i].Value,
					}
					if c[i].NextToken != nil {
						next := *c[i].NextToken
						out.NextToken = aws.String(next)
					}
					req.Data = out
				})
				return r, nil
			},
		}

		for p.Next() {
			data := p.Page().(*testPageOutput)

			if e, a := aws.StringValue(c[i].Value), aws.StringValue(data.Value); e != a {
				t.Errorf("%d, expect Value to be %q, got %q", i, e, a)
			}
			if e, a := aws.StringValue(c[i].NextToken), aws.StringValue(data.NextToken); e != a {
				t.Errorf("%d, expect NextToken to be %q, got %q", i, e, a)
			}

			i++
		}
		if e, a := len(c), i; e != a {
			t.Errorf("expected to process %d pages, did %d", e, a)
		}
		if err := p.Err(); err != nil {
			t.Fatalf("%d, expected no error, got %v", i, err)
		}
	}
}
