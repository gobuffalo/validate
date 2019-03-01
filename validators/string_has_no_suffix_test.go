package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasNoSuffix(t *testing.T) {

	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "t", false},
		{"test_fail", "zzz", true},
		{" test with space", " zzz", true},
		{" test with space second", "second", false},
	}

	for _, testCase := range cases {
		v := StringHasNoSuffix{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	for _, testCase := range cases {
		v := StringHasNoSuffix{Name: "strings1", Field: testCase.str1, ComparedField: testCase.str2, ComparedName: "strings2"}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
		if !testCase.expected {
			r.Contains(e.Get("strings1"), "'strings1' ends with content of 'strings2'")
		}
	}
}
