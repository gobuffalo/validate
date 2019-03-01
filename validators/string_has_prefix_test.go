package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasPrefix(t *testing.T) {

	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "t", true},
		{"test_fail", "zzz", false},
		{" test with space", " test", true},
		{" test with space second", "test", false},
	}

	for _, testCase := range cases {
		v := StringHasPrefix{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	for _, testCase := range cases {
		v := StringHasPrefix{Name: "strings1", Field: testCase.str1, ComparedField: testCase.str2, ComparedName: "strings2"}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
		if !testCase.expected {
			r.Contains(e.Get("strings1"), "'strings1' does not start with content of 'strings2'")
		}
	}

	v := StringHasPrefix{Name: "strings", Field: "test_fail", ComparedField: "zzz"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"'test_fail' does not start with 'zzz'"}, e.Get("strings"))
}
