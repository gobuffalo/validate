package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringsMatch_Validate(t *testing.T) {

	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "test", true},
		{"test_fail", "test_true", false},
		{"test with space", " test with space ", true},
		{" test with space second", " test with space second       ", true},
	}

	for _, testCase := range cases {
		v := StringsMatch{Name: "strings", Field: testCase.str1, Field2: testCase.str2}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	v := StringsMatch{Name: "strings", Field: "test_fail", Field2: "test", Message: "String doesn't match"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"String doesn't match"}, e.Get("strings"))
}

func BenchmarkStringsMatch_IsValid_Valid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringsMatch{Name: "strings", Field: " Some string ", Field2: " Some string "}
		v.Validate(e)
	}
}

func BenchmarkStringsMatch_IsValid_InValid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringsMatch{Name: "strings", Field: " Some string ", Field2: " Some string failure"}
		v.Validate(e)
	}
}
