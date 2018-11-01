package pagin8

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/awstesting"
	"github.com/aws/aws-sdk-go/awstesting/unit"
)

func TestMockPagesTruncation(t *testing.T) {
	svc := New(unit.Session)
	tokens, pages, numPages, gotToEnd := []*string{}, []*int64{}, 0, false

	reqNum := 0
	resps := []*MockOutput{
		{
			NextToken:   aws.String("a"),
			IsTruncated: aws.Bool(true),
			List: []*int64{
				aws.Int64(0),
				aws.Int64(1),
				aws.Int64(2),
			},
		},
		{
			NextToken:   aws.String("b"),
			IsTruncated: aws.Bool(true),
			List: []*int64{
				aws.Int64(3),
			},
		},
		{
			IsTruncated: aws.Bool(false),
			NextToken:   aws.String("c"),
			List: []*int64{
				aws.Int64(6),
				aws.Int64(7),
				aws.Int64(8),
				aws.Int64(9),
			},
		},
	}

	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()
	svc.Handlers.Build.RemoveByName("crr.endpointdiscovery")
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		in := r.Params.(*MockInput)
		if in.Token == nil {
			tokens = append(tokens, nil)
		} else if len(*in.Token) != 0 {
			tokens = append(tokens, in.Token)
		}
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		r.Data = resps[reqNum]
		reqNum++
	})

	params := &MockInput{}

	err := svc.MockPages(params, func(p *MockOutput, last bool) bool {
		numPages++
		for _, elem := range p.List {
			pages = append(pages, elem)
		}

		if last {
			if gotToEnd {
				t.Errorf("Got to end occurred twice. Should have only occurred once")
			}
			gotToEnd = true
		}
		return true
	})

	if err != nil {
		t.Errorf("expected no error, but received %v", err)
	}

	expectedTokens := []*string{
		nil,
		aws.String("a"),
		aws.String("b"),
	}
	if e, a := expectedTokens, tokens; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}

	expectedPages := []*int64{
		aws.Int64(0),
		aws.Int64(1),
		aws.Int64(2),
		aws.Int64(3),
		aws.Int64(6),
		aws.Int64(7),
		aws.Int64(8),
		aws.Int64(9),
	}
	if e, a := expectedPages, pages; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if e, a := 3, numPages; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if e, a := true, gotToEnd; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}
}

func TestPaginationEachPage(t *testing.T) {
	svc := New(unit.Session)
	tokens, pages, numPages, gotToEnd := []string{}, []int64{}, 0, false

	reqNum := 0
	resps := []*MockOutput{
		{List: []*int64{aws.Int64(0), aws.Int64(1)}, NextToken: aws.String("a"), IsTruncated: aws.Bool(true)},
		{List: []*int64{aws.Int64(2), aws.Int64(3), aws.Int64(4)}, NextToken: aws.String("b"), IsTruncated: aws.Bool(true)},
		{List: []*int64{aws.Int64(5)}, IsTruncated: aws.Bool(false)},
	}

	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()
	svc.Handlers.Build.RemoveByName("crr.endpointdiscovery")
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		in := r.Params.(*MockInput)
		if in.Token == nil {
			tokens = append(tokens, "")
		} else {
			tokens = append(tokens, *in.Token)
		}
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		r.Data = resps[reqNum]
		reqNum++
	})

	params := &MockInput{}
	req, _ := svc.MockRequest(params)
	err := req.EachPage(func(p interface{}, last bool) bool {
		numPages++
		for _, t := range p.(*MockOutput).List {
			pages = append(pages, *t)
		}
		if last {
			if gotToEnd {
				t.Errorf("Pagination has already ended. Unexpected additional request")
			}
			gotToEnd = true
		}

		return true
	})

	if e, a := 3, numPages; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if e, a := true, gotToEnd; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if err != nil {
		t.Errorf("unexpected error occurred %v", err)
	}

	expectedTokens := []string{"", "a", "b"}
	if e, a := expectedTokens, tokens; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}

	expectedPages := []int64{0, 1, 2, 3, 4, 5}
	if e, a := expectedPages, pages; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}
}

func TestPaginationEarlyExit(t *testing.T) {
	svc := New(unit.Session)
	tokens, pages, numPages, gotToEnd := []string{}, []int64{}, 0, false

	reqNum := 0
	resps := []*MockOutput{
		{List: []*int64{aws.Int64(0), aws.Int64(1)}, NextToken: aws.String("a"), IsTruncated: aws.Bool(true)},
		{List: []*int64{aws.Int64(2), aws.Int64(3), aws.Int64(4)}, NextToken: aws.String("b"), IsTruncated: aws.Bool(true)},
		{List: []*int64{aws.Int64(5)}, IsTruncated: aws.Bool(false)},
	}

	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()
	svc.Handlers.Build.RemoveByName("crr.endpointdiscovery")
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		in := r.Params.(*MockInput)
		if in.Token == nil {
			tokens = append(tokens, "")
		} else {
			tokens = append(tokens, *in.Token)
		}
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		r.Data = resps[reqNum]
		reqNum++
	})

	params := &MockInput{}
	req, _ := svc.MockRequest(params)
	err := req.EachPage(func(p interface{}, last bool) bool {
		numPages++
		if numPages == 2 {
			return false
		}

		for _, t := range p.(*MockOutput).List {
			pages = append(pages, *t)
		}
		if last {
			if gotToEnd {
				t.Errorf("Pagination has already ended. Unexpected additional request")
			}
			gotToEnd = true
		}

		return true
	})

	if e, a := 2, numPages; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if e, a := false, gotToEnd; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if err != nil {
		t.Errorf("unexpected error occurred %v", err)
	}

	expectedTokens := []string{"", "a"}
	if e, a := expectedTokens, tokens; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}

	expectedPages := []int64{0, 1}
	if e, a := expectedPages, pages; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}
}

func TestPaginationSkip(t *testing.T) {
	svc := New(unit.Session)
	tokens, pages, numPages, gotToEnd := []string{}, []int64{}, 0, false

	reqNum := 0
	resps := []*MockOutput{
		{},
	}

	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()
	svc.Handlers.Build.RemoveByName("crr.endpointdiscovery")
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		in := r.Params.(*MockInput)
		if in.Token == nil {
			tokens = append(tokens, "")
		} else {
			tokens = append(tokens, *in.Token)
		}
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		r.Data = resps[reqNum]
		reqNum++
	})

	params := &MockInput{}
	req, _ := svc.MockRequest(params)
	err := req.EachPage(func(p interface{}, last bool) bool {
		numPages++
		for _, t := range p.(*MockOutput).List {
			pages = append(pages, *t)
		}
		if last {
			if gotToEnd {
				t.Errorf("Pagination has already ended. Unexpected additional request")
			}
			gotToEnd = true
		}

		return true
	})

	if e, a := 1, numPages; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if e, a := true, gotToEnd; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if err != nil {
		t.Errorf("unexpected error occurred %v", err)
	}

	expectedTokens := []string{""}
	if e, a := expectedTokens, tokens; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}

	expectedPages := []int64{}
	if e, a := expectedPages, pages; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}
}

func TestPaginationNilToken(t *testing.T) {
	svc := New(unit.Session)
	tokens, pages, numPages, gotToEnd := []*string{}, []*int64{}, 0, false

	reqNum := 0
	resps := []*MockOutput{
		{
			IsTruncated: aws.Bool(true),
			List: []*int64{
				aws.Int64(0),
				aws.Int64(1),
				aws.Int64(2),
			},
		},
		{
			NextToken:   aws.String("b"),
			IsTruncated: aws.Bool(true),
			List: []*int64{
				aws.Int64(3),
			},
		},
		{
			IsTruncated: aws.Bool(false),
			NextToken:   aws.String("c"),
			List: []*int64{
				aws.Int64(6),
				aws.Int64(7),
				aws.Int64(8),
				aws.Int64(9),
			},
		},
	}

	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()
	svc.Handlers.Build.RemoveByName("crr.endpointdiscovery")
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		in := r.Params.(*MockInput)
		if in.Token == nil {
			tokens = append(tokens, nil)
		} else if len(*in.Token) != 0 {
			tokens = append(tokens, in.Token)
		}
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		r.Data = resps[reqNum]
		reqNum++
	})

	params := &MockInput{}

	err := svc.MockPages(params, func(p *MockOutput, last bool) bool {
		numPages++
		for _, elem := range p.List {
			pages = append(pages, elem)
		}

		if last {
			if gotToEnd {
				t.Errorf("Got to end occurred twice. Should have only occurred once")
			}
			gotToEnd = true
		}
		return true
	})

	if err != nil {
		t.Errorf("expected no error, but received %v", err)
	}

	expectedTokens := []*string{nil}
	if e, a := expectedTokens, tokens; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}

	expectedPages := []*int64{
		aws.Int64(0),
		aws.Int64(1),
		aws.Int64(2),
	}
	if e, a := expectedPages, pages; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if e, a := 1, numPages; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if e, a := true, gotToEnd; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}
}

func TestPaginationNilInput(t *testing.T) {
	svc := New(unit.Session)
	tokens, pages, numPages, gotToEnd := []*string{}, []*int64{}, 0, false

	reqNum := 0
	resps := []*MockOutput{
		{
			NextToken:   aws.String("a"),
			IsTruncated: aws.Bool(true),
			List: []*int64{
				aws.Int64(0),
				aws.Int64(1),
				aws.Int64(2),
			},
		},
		{
			NextToken:   aws.String("b"),
			IsTruncated: aws.Bool(true),
			List: []*int64{
				aws.Int64(3),
			},
		},
		{
			IsTruncated: aws.Bool(false),
			NextToken:   aws.String("c"),
			List: []*int64{
				aws.Int64(6),
				aws.Int64(7),
				aws.Int64(8),
				aws.Int64(9),
			},
		},
	}

	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.ValidateResponse.Clear()
	svc.Handlers.Build.RemoveByName("crr.endpointdiscovery")
	svc.Handlers.Build.PushBack(func(r *request.Request) {
		in := r.Params.(*MockInput)
		if in.Token == nil {
			tokens = append(tokens, nil)
		} else if len(*in.Token) != 0 {
			tokens = append(tokens, in.Token)
		}
	})
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		r.Data = resps[reqNum]
		reqNum++
	})

	err := svc.MockPages(nil, func(p *MockOutput, last bool) bool {
		numPages++
		for _, elem := range p.List {
			pages = append(pages, elem)
		}

		if last {
			if gotToEnd {
				t.Errorf("Got to end occurred twice. Should have only occurred once")
			}
			gotToEnd = true
		}
		return true
	})

	if err != nil {
		t.Errorf("expected no error, but received %v", err)
	}

	expectedTokens := []*string{
		nil,
		aws.String("a"),
		aws.String("b"),
	}
	if e, a := expectedTokens, tokens; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}

	expectedPages := []*int64{
		aws.Int64(0),
		aws.Int64(1),
		aws.Int64(2),
		aws.Int64(3),
		aws.Int64(6),
		aws.Int64(7),
		aws.Int64(8),
		aws.Int64(9),
	}
	if e, a := expectedPages, pages; !reflect.DeepEqual(e, a) {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if e, a := 3, numPages; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}

	if e, a := true, gotToEnd; e != a {
		t.Errorf("expected %v, but received %v", e, a)
	}
}

func TestPaginationWithContextNilInput(t *testing.T) {
	// Code generation doesn't have a great way to verify the code is correct
	// other than being run via unit tests in the SDK. This should be fixed
	// So code generation can be validated independently.

	client := New(unit.Session)
	client.Handlers.Validate.Clear()
	client.Handlers.Send.Clear() // mock sending
	client.Handlers.Unmarshal.Clear()
	client.Handlers.UnmarshalMeta.Clear()
	client.Handlers.ValidateResponse.Clear()
	client.Handlers.Build.RemoveByName("crr.endpointdiscovery")
	client.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		r.Data = &MockOutput{}
	})

	gotToEnd := false
	numPages := 0
	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	err := client.MockPagesWithContext(ctx, nil, func(p *MockOutput, last bool) bool {
		numPages++
		if last {
			gotToEnd = true
		}
		return true
	})

	if err != nil {
		t.Fatalf("expect no error, but got %v", err)
	}
	if e, a := 1, numPages; e != a {
		t.Errorf("expect %d number pages but got %d", e, a)
	}
	if !gotToEnd {
		t.Errorf("expect to of gotten to end, did not")
	}
}

// Benchmarks
var benchResps = []*MockOutput{
	{NextToken: aws.String("a"), IsTruncated: aws.Bool(true), List: []*int64{aws.Int64(0), aws.Int64(1), aws.Int64(2)}},
	{NextToken: aws.String("b"), IsTruncated: aws.Bool(true), List: []*int64{aws.Int64(3)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(true), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
	{IsTruncated: aws.Bool(false), NextToken: aws.String("c"), List: []*int64{aws.Int64(6), aws.Int64(7), aws.Int64(8), aws.Int64(9)}},
}

var benchClient = func() *Pagin8 {
	svc := New(unit.Session)
	svc.Handlers.Send.Clear() // mock sending
	svc.Handlers.Unmarshal.Clear()
	svc.Handlers.UnmarshalMeta.Clear()
	svc.Handlers.Build.RemoveByName("crr.endpointdiscovery")
	svc.Handlers.ValidateResponse.Clear()
	return svc
}

func BenchmarkCodegenIterator(b *testing.B) {
	reqNum := 0
	svc := benchClient()
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		r.Data = benchResps[reqNum]
		reqNum++
	})

	input := &MockInput{}
	iter := func(fn func(*MockOutput, bool) bool) error {
		page, _ := svc.MockRequest(input)
		for ; page != nil; page = page.NextPage() {
			page.Send()
			out := page.Data.(*MockOutput)
			if result := fn(out, !page.HasNextPage()); page.Error != nil || !result {
				return page.Error
			}
		}
		return nil
	}

	for i := 0; i < b.N; i++ {
		reqNum = 0
		iter(func(p *MockOutput, last bool) bool {
			return true
		})
	}
}

func BenchmarkEachPageIterator(b *testing.B) {
	reqNum := 0
	svc := benchClient()
	svc.Handlers.Unmarshal.PushBack(func(r *request.Request) {
		r.Data = benchResps[reqNum]
		reqNum++
	})

	input := &MockInput{}
	for i := 0; i < b.N; i++ {
		reqNum = 0
		req, _ := svc.MockRequest(input)
		req.EachPage(func(p interface{}, last bool) bool {
			return true
		})
	}
}
