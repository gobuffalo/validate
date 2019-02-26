package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringsAreEqual(t *testing.T) {

	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "test", true},
		{"test_fail", "test_true", false},
		{"test with space", " test with space ", false},
		{" test with space second", " test with space second       ", false},
	}

	for _, testCase := range cases {
		v := StringsAreEqual{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	v := StringsAreEqual{Name: "strings", Field: "test_fail", ComparedField: "test", Message: "String doesn't match"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"String doesn't match"}, e.Get("strings"))
}

func Test_StringsAreIEqual(t *testing.T) {

	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "TesT", true},
		{"test_fail", "Test_truE", false},
		{"test with space", " Test with spacE ", false},
		{" test with space second", " Test with space seconD       ", false},
	}

	for _, testCase := range cases {
		v := StringsAreEqual{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2, CaseInsensitive: true}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	v := StringsAreEqual{Name: "strings", Field: "test_fail", ComparedField: "test", CaseInsensitive: true, Message: "String doesn't match"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"String doesn't match"}, e.Get("strings"))
}

func BenchmarkStringsAreEqual_Valid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringsAreEqual{Name: "strings", Field: " Some string ", ComparedField: " Some string "}
		v.Validate(e)
	}
}

func BenchmarkStringsAreEqual_InValid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringsAreEqual{Name: "strings", Field: " Some string ", ComparedField: " Some string failure"}
		v.Validate(e)
	}
}
