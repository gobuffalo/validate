package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringsAreNotEqual(t *testing.T) {

	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "test", false},
		{"test_fail", "test_true", true},
		{"test with space", " test with space ", true},
		{" test with space second", " test with space second       ", true},
	}

	for _, testCase := range cases {
		v := StringsAreNotEqual{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	for _, testCase := range cases {
		v := StringsAreNotEqual{Name: "strings1", Field: testCase.str1, ComparedField: testCase.str2, ComparedName: "strings2"}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
		if !testCase.expected {
			r.Contains(e.Get("strings1"), "strings1 is equal to strings2")
		}
	}

	v := StringsAreNotEqual{Name: "strings", Field: "test", ComparedField: "test"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"test is equal to test"}, e.Get("strings"))
}

func Test_StringsAreNotIEqual(t *testing.T) {

	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "TesT", false},
		{"test_fail", "Test_truE", true},
		{"test with space", " Test with spacE ", true},
		{" test with space second", " Test with space seconD       ", true},
	}

	for _, testCase := range cases {
		v := StringsAreNotEqual{Name: "strings", Field: testCase.str1, ComparedField: testCase.str2, CaseInsensitive: true}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
	}

	for _, testCase := range cases {
		v := StringsAreNotEqual{Name: "strings1", Field: testCase.str1, ComparedField: testCase.str2, ComparedName: "strings2", CaseInsensitive: true}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)
		if !testCase.expected {
			r.Contains(e.Get("strings1"), "strings1 is iequal to strings2")
		}
	}

	v := StringsAreNotEqual{Name: "strings", Field: "test", ComparedField: "test", CaseInsensitive: true}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"test is iequal to test"}, e.Get("strings"))
}

func BenchmarkStringsAreNotEqual_Valid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringsAreNotEqual{Name: "strings", Field: " Some string ", ComparedField: " Some string "}
		v.Validate(e)
	}
}

func BenchmarkStringsAreNotEqual_InValid(b *testing.B) {
	e := validator.NewErrors()
	for i := 0; i <= b.N; i++ {
		v := StringsAreNotEqual{Name: "strings", Field: " Some string ", ComparedField: " Some string failure"}
		v.Validate(e)
	}
}
