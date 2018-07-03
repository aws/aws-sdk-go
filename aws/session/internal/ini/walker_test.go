package ini

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestDataFiles(t *testing.T) {
	cases := []struct {
		path               string
		expectedParseError bool
		expectedWalkError  bool

		expectedPath string
	}{
		{
			path:         "./testdata/valid/empty_profile",
			expectedPath: "./testdata/valid/empty_profile_expected",
		},
		{
			path:         "./testdata/valid/simple_profile",
			expectedPath: "./testdata/valid/simple_profile_expected",
		},
		{
			path:         "./testdata/valid/nested_profile",
			expectedPath: "./testdata/valid/nested_profile_expected",
		},
		{
			path:         "./testdata/valid/commented_profile",
			expectedPath: "./testdata/valid/commented_profile_expected",
		},
		{
			path:         "./testdata/valid/sections_profile",
			expectedPath: "./testdata/valid/sections_profile_expected",
		},
		{
			path:         "./testdata/valid/number_lhs_expr",
			expectedPath: "./testdata/valid/number_lhs_expr_expected",
		},
		{
			path:               "./testdata/invalid/incomplete_section_profile",
			expectedParseError: true,
		},
		{
			path:               "./testdata/invalid/syntax_error_comment",
			expectedParseError: true,
		},
		{
			path:               "./testdata/invalid/bad_syntax_1",
			expectedParseError: true,
		},
		{
			path:              "./testdata/invalid/bad_syntax_2",
			expectedWalkError: true,
		},
	}

	for _, c := range cases {
		f, err := os.Open(c.path)
		if err != nil {
			t.Errorf("unexpected error, %v", err)
		}

		tree, err := Parse(f)
		if err != nil && !c.expectedParseError {
			t.Errorf("unexpected error, %v", err)
		} else if err == nil && c.expectedParseError {
			t.Errorf("expected error, but received none")
		}

		if c.expectedParseError {
			continue
		}

		v := NewSharedConfigVisitor()
		err = Walk(tree, v)
		if err != nil && !c.expectedWalkError {
			t.Errorf("unexpected error, %v", err)
		} else if err == nil && c.expectedWalkError {
			t.Errorf("expected error, but received none")
		}

		if len(c.expectedPath) == 0 {
			continue
		}

		e := map[string]interface{}{}

		b, err := ioutil.ReadFile(c.expectedPath)
		if err != nil {
			t.Errorf("unexpected error opening expected file, %v", err)
		}

		err = json.Unmarshal(b, &e)
		if err != nil {
			t.Errorf("unexpected error during deserialization, %v", err)
		}

		for profile, tableIface := range e {
			p, ok := v.Tables.GetSection(profile)
			if !ok {
				t.Fatal("could not find profile " + profile)
			}

			table := tableIface.(map[string]interface{})
			for k, v := range table {
				switch e := v.(type) {
				case string:
					a := p.String(k)
					if e != a {
						t.Errorf("expected %v, but received %v", e, a)
					}
				case int:
					a := p.Int(k)
					if e != a {
						t.Errorf("expected %v, but received %v", e, a)
					}
				case float64:
					a := p.Float64(k)
					if e != a {
						t.Errorf("expected %v, but received %v", e, a)
					}
				default:
					t.Errorf("unexpected type: %T", e)
				}
			}
		}

		f.Close()
	}
}
