// +build go1.7

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
	}{
		{
			path: "./testdata/valid/empty_profile",
		},
		{
			path: "./testdata/valid/array_profile",
		},
		{
			path: "./testdata/valid/simple_profile",
		},
		{
			path: "./testdata/valid/arn_profile",
		},
		{
			path: "./testdata/valid/commented_profile",
		},
		{
			path: "./testdata/valid/sections_profile",
		},
		{
			path: "./testdata/valid/number_lhs_expr",
		},
		{
			path: "./testdata/valid/base_numbers_profile",
		},
		{
			path: "./testdata/valid/exponent_profile",
		},
		{
			path: "./testdata/valid/escaped_profile",
		},
		{
			path: "./testdata/valid/global_values_profile",
		},
		{
			path: "./testdata/valid/utf_8_profile",
		},
		{
			path: "./testdata/valid/profile_name",
		},
		{
			path:               "./testdata/invalid/bad_syntax_1",
			expectedParseError: true,
		},
		{
			path:               "./testdata/invalid/incomplete_section_profile",
			expectedParseError: true,
		},
		{
			path:               "./testdata/invalid/syntax_error_comment",
			expectedParseError: true,
		},
	}

	for i, c := range cases {
		t.Run(c.path, func(t *testing.T) {
			f, err := os.Open(c.path)
			if err != nil {
				t.Errorf("unexpected error, %v", err)
			}

			tree, err := ParseAST(f)
			if err != nil && !c.expectedParseError {
				t.Errorf("%d: unexpected error, %v", i+1, err)
			} else if err == nil && c.expectedParseError {
				t.Errorf("%d: expected error, but received none", i+1)
			}

			if c.expectedParseError {
				return
			}

			v := NewDefaultVisitor()
			err = Walk(tree, v)
			if err != nil && !c.expectedWalkError {
				t.Errorf("%d: unexpected error, %v", i+1, err)
			} else if err == nil && c.expectedWalkError {
				t.Errorf("%d: expected error, but received none", i+1)
			}

			expectedPath := c.path + "_expected"
			e := map[string]interface{}{}

			b, err := ioutil.ReadFile(expectedPath)
			if err != nil {
				return
			}

			err = json.Unmarshal(b, &e)
			if err != nil {
				t.Errorf("unexpected error during deserialization, %v", err)
			}

			for profile, tableIface := range e {
				p, ok := v.Sections.GetSection(profile)
				if !ok {
					t.Fatal("could not find profile " + profile)
				}

				table := tableIface.(map[string]interface{})
				for k, v := range table {
					switch e := v.(type) {
					case string:
						a := p.String(k)
						if e != a {
							t.Errorf("%d: expected %v, but received %v", i+1, e, a)
						}
					case int:
						a := p.Int(k)
						if int64(e) != a {
							t.Errorf("%d: expected %v, but received %v", i+1, e, a)
						}
					case float64:
						v := p.values[k]
						if v.Type == IntegerType {
							a := p.Int(k)
							if int64(e) != a {
								t.Errorf("%d: expected %v, but received %v", i+1, e, a)
							}
						} else {
							a := p.Float64(k)
							if e != a {
								t.Errorf("%d: expected %v, but received %v", i+1, e, a)
							}
						}
					default:
						t.Errorf("unexpected type: %T", e)
					}
				}
			}

			f.Close()
		})
	}
}
