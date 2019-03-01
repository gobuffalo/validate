package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasSuffix(t *testing.T) {

	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "t", true},
		{"test_fail", "zzz", false},
		{" test with space", " zzz", false},
		{" test with space second", "second", true},
	}

	for _, testCase := range cases {
		v := StringHasSuffix{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	for _, testCase := range cases {
		v := StringHasSuffix{Name: "strings1", Field: testCase.str1, ComparedField: testCase.str2, ComparedName: "strings2"}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
		if !testCase.expected {
			r.Contains(e.Get("strings1"), "'strings1' does not end with content of 'strings2'")
		}
	}
}
