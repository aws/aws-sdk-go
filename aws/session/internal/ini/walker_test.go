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
			path:         "./testdata/valid/nested_profile_2",
			expectedPath: "./testdata/valid/nested_profile_2_expected",
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
			path:         "./testdata/valid/base_numbers_profile",
			expectedPath: "./testdata/valid/base_numbers_profile_expected",
		},
		{
			path:         "./testdata/valid/exponent_profile",
			expectedPath: "./testdata/valid/exponent_profile_expected",
		},
		{
			path:         "./testdata/valid/escaped_profile",
			expectedPath: "./testdata/valid/escaped_profile_expected",
		},
		{
			path:         "./testdata/valid/global_values_profile",
			expectedPath: "./testdata/valid/global_values_profile_expected",
		},
		{
			path:              "./testdata/invalid/bad_syntax_1",
			expectedWalkError: true,
		},
		{
			path:              "./testdata/invalid/incomplete_section_profile",
			expectedWalkError: true,
		},
		{
			path:              "./testdata/invalid/syntax_error_comment",
			expectedWalkError: true,
		},
	}

	for i, c := range cases {
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
			continue
		}

		v := NewSharedConfigVisitor()
		err = Walk(tree, v)
		if err != nil && !c.expectedWalkError {
			t.Errorf("%d: unexpected error, %v", i+1, err)
		} else if err == nil && c.expectedWalkError {
			t.Errorf("%d: expected error, but received none", i+1)
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
					tok := p.Values[k]
					v := tok.(literalToken)
					if v.Value.Type == IntegerType {
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
	}
}
