package validator

import (
	"encoding/json"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type a struct{}

func (s *a) Validate(e *Errors) {
	e.Add(
		"a",
		"a",
	)
}

type b struct{}

func (s *b) Validate(e *Errors) {
	e.Add(
		"b",
		"[0]b",
	)
	e.Add(
		"b",
		"b",
	)
	e.Add(
		"b.b",
		"[0]b.b",
	)
	e.Add(
		"b.b",
		"b.b",
	)
	e.Add(
		"bb.bb",
		"bb.bb",
	)
	e.Add(
		"bb",
		"bb",
	)
	e.Add(
		"bbb[1]",
		"bbb[1]",
	)
	e.Add(
		"bbb[1].bb[3]",
		"bbb[1].bb[3]",
	)
	e.Add(
		"bbbb[5].bbbb",
		"bbbb[5].bbbb",
	)
}

type c struct{}

func (s *c) Validate(e *Errors) {
	e.Add(
		"d.d",
		"d",
	)
}

type z struct{}

func (s *z) Validate(e *Errors) {
}

func TestCount(t *testing.T) {

	r := require.New(t)

	e := Validate(&a{}, &b{})
	r.Equal(10, e.Count())

	e = Validate(&z{})
	r.Equal(0, e.Count())
}

func TestCountP(t *testing.T) {

	r := require.New(t)

	e := ValidateP(&a{}, &b{})
	r.Equal(10, e.Count())

	e = ValidateP(&z{})
	r.Equal(0, e.Count())
}

func TestHasAny(t *testing.T) {

	r := require.New(t)

	e := Validate(&a{})
	r.Equal(true, e.HasAny())

	e = Validate(&z{})
	r.Equal(false, e.HasAny())
}

func TestHasAnyP(t *testing.T) {

	r := require.New(t)

	e := ValidateP(&a{})
	r.Equal(true, e.HasAny())

	e = ValidateP(&z{})
	r.Equal(false, e.HasAny())
}

func TestValidate(t *testing.T) {

	r := require.New(t)

	e := Validate(&a{}, &b{})
	r.Error(e)

	e = Validate(&z{})
	r.Nil(e)
}

func TestValidateP(t *testing.T) {

	r := require.New(t)

	e := ValidateP(&a{}, &b{})
	r.Error(e)

	e = ValidateP(&z{})
	r.Nil(e)
}

func TestExists(t *testing.T) {

	r := require.New(t)

	e := Validate(&a{}, &b{})

	r.Equal(true, e.Exists("a"))
	r.Equal(true, e.Exists("b"))
	r.Equal(true, e.Exists("b.b"))
	r.Equal(true, e.Exists("bb.bb"))
	r.Equal(true, e.Exists("bb"))
	r.Equal(true, e.Exists("bbb[1]"))
	r.Equal(true, e.Exists("bbb[1].bb[3]"))
	r.Equal(true, e.Exists("bbbb[5].bbbb"))
	r.Equal(false, e.Exists("z"))
	r.Equal(false, e.Exists("z.z"))
	r.Equal(false, e.Exists("z.[5]z"))
	r.Equal(false, e.Exists("z[0].z"))

	e = Validate(&z{})
	r.Equal(false, e.Exists("a"))
}

func TestExistsP(t *testing.T) {

	r := require.New(t)

	e := ValidateP(&a{}, &b{})

	r.Equal(true, e.Exists("a"))
	r.Equal(true, e.Exists("b"))
	r.Equal(true, e.Exists("b.b"))
	r.Equal(true, e.Exists("bb.bb"))
	r.Equal(true, e.Exists("bb"))
	r.Equal(true, e.Exists("bbb[1]"))
	r.Equal(true, e.Exists("bbb[1].bb[3]"))
	r.Equal(true, e.Exists("bbbb[5].bbbb"))
	r.Equal(false, e.Exists("z"))
	r.Equal(false, e.Exists("z.z"))
	r.Equal(false, e.Exists("z.[5]z"))
	r.Equal(false, e.Exists("z[0].z"))

	e = ValidateP(&z{})
	r.Equal(false, e.Exists("a"))
}

func TestGet(t *testing.T) {

	r := require.New(t)

	e := Validate(&a{}, &b{})

	r.Equal(1, len(e.Get("a")))
	r.Contains(e.Get("a"), "a")

	r.Equal(2, len(e.Get("b")))
	r.Contains(e.Get("b"), "[0]b")

	r.Equal(2, len(e.Get("b.b")))
	r.Contains(e.Get("b.b"), "[0]b.b")

	r.Equal(1, len(e.Get("bb.bb")))
	r.Contains(e.Get("bb.bb"), "bb.bb")

	r.Equal(1, len(e.Get("bb")))
	r.Contains(e.Get("bb"), "bb")

	r.Equal(1, len(e.Get("bbb[1]")))
	r.Contains(e.Get("bbb[1]"), "bbb[1]")

	r.Equal(1, len(e.Get("bbb[1].bb[3]")))
	r.Contains(e.Get("bbb[1].bb[3]"), "bbb[1].bb[3]")

	r.Equal(1, len(e.Get("bbbb[5].bbbb")))
	r.Contains(e.Get("bbbb[5].bbbb"), "bbbb[5].bbbb")

	r.Equal(0, len(e.Get("z")))
	r.Equal(0, len(e.Get("z.z")))
	r.Equal(0, len(e.Get("z.[5]z")))
	r.Equal(0, len(e.Get("z[0].z")))

	e = Validate(&z{})
	r.Equal(0, len(e.Get("a")))
}

func TestGetP(t *testing.T) {

	r := require.New(t)

	e := ValidateP(&a{}, &b{})

	r.Equal(1, len(e.Get("a")))
	r.Contains(e.Get("a"), "a")

	r.Equal(2, len(e.Get("b")))
	r.Contains(e.Get("b"), "[0]b")

	r.Equal(2, len(e.Get("b.b")))
	r.Contains(e.Get("b.b"), "[0]b.b")

	r.Equal(1, len(e.Get("bb.bb")))
	r.Contains(e.Get("bb.bb"), "bb.bb")

	r.Equal(1, len(e.Get("bb")))
	r.Contains(e.Get("bb"), "bb")

	r.Equal(1, len(e.Get("bbb[1]")))
	r.Contains(e.Get("bbb[1]"), "bbb[1]")

	r.Equal(1, len(e.Get("bbb[1].bb[3]")))
	r.Contains(e.Get("bbb[1].bb[3]"), "bbb[1].bb[3]")

	r.Equal(1, len(e.Get("bbbb[5].bbbb")))
	r.Contains(e.Get("bbbb[5].bbbb"), "bbbb[5].bbbb")

	r.Equal(0, len(e.Get("z")))
	r.Equal(0, len(e.Get("z.z")))
	r.Equal(0, len(e.Get("z.[5]z")))
	r.Equal(0, len(e.Get("z[0].z")))

	e = Validate(&z{})
	r.Equal(0, len(e.Get("a")))
}

func TestKeys(t *testing.T) {

	r := require.New(t)

	e := Validate(&a{}, &b{})

	r.Contains(e.Keys(), "a")
	r.Contains(e.Keys(), "b")
	r.Contains(e.Keys(), "b.b")
	r.Contains(e.Keys(), "bb.bb")
	r.Contains(e.Keys(), "bb")
	r.Contains(e.Keys(), "bbb[1]")
	r.Contains(e.Keys(), "bbb[1].bb[3]")
	r.Contains(e.Keys(), "bbbb[5].bbbb")
	r.NotContains(e.Keys(), "z")
	r.NotContains(e.Keys(), "z.z")
	r.NotContains(e.Keys(), "z.[5]z")
	r.NotContains(e.Keys(), "z[0].z")

	// test correct sort order of keys
	k := e.Keys()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(k), func(i, j int) { k[i], k[j] = k[j], k[i] })

	// sort keys in reverse order
	sort.Sort(sort.Reverse(sort.StringSlice(k)))

	r.Exactlyf(k, e.Keys(), "keys must be sorted in descending order")

	e = Validate(&z{})
	r.NotContains(e.Keys(), "a")
}

func TestKeysP(t *testing.T) {

	r := require.New(t)

	e := ValidateP(&a{}, &b{})

	r.Contains(e.Keys(), "a")
	r.Contains(e.Keys(), "b")
	r.Contains(e.Keys(), "b.b")
	r.Contains(e.Keys(), "bb.bb")
	r.Contains(e.Keys(), "bb")
	r.Contains(e.Keys(), "bbb[1]")
	r.Contains(e.Keys(), "bbb[1].bb[3]")
	r.Contains(e.Keys(), "bbbb[5].bbbb")
	r.NotContains(e.Keys(), "z")
	r.NotContains(e.Keys(), "z.z")
	r.NotContains(e.Keys(), "z.[5]z")
	r.NotContains(e.Keys(), "z[0].z")

	// test correct sort order of keys
	k := e.Keys()
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(k), func(i, j int) { k[i], k[j] = k[j], k[i] })

	// sort keys in reverse order
	sort.Sort(sort.Reverse(sort.StringSlice(k)))

	r.Exactlyf(k, e.Keys(), "keys must be sorted in descending order")

	e = ValidateP(&z{})
	r.NotContains(e.Keys(), "a")
}

func TestData(t *testing.T) {

	r := require.New(t)

	expectedType := make(map[string]interface{})

	e := Validate(&a{}, &b{})
	r.IsType(expectedType, e.Data())

	e = Validate(&z{})
	r.Nil(e)
}

func TestDataP(t *testing.T) {

	r := require.New(t)

	expectedType := make(map[string]interface{})

	e := ValidateP(&a{}, &b{})
	r.IsType(expectedType, e.Data())

	e = ValidateP(&z{})
	r.Nil(e)
}

func TestString(t *testing.T) {

	r := require.New(t)

	e := Validate(&c{})

	b, err := json.Marshal(e.Data())
	r.Nil(err)

	r.JSONEq(`{"d.d":["d"]}`, e.JSON())
	r.Exactly(string(b), e.JSON())

	e = Validate(&z{})
	r.Nil(e)
}

func TestStringP(t *testing.T) {

	r := require.New(t)

	e := ValidateP(&c{})

	b, err := json.Marshal(e.Data())
	r.Nil(err)

	r.JSONEq(`{"d":{"d":["d"]}}`, e.JSON())
	r.Exactly(string(b), e.JSON())

	e = ValidateP(&z{})
	r.Nil(e)
}

func TestError(t *testing.T) {

	r := require.New(t)

	e := Validate(&c{})

	r.Exactly(e.JSON(), e.Error())

	e = Validate(&z{})
	r.Nil(e)
}

func TestErrorP(t *testing.T) {

	r := require.New(t)

	e := ValidateP(&c{})

	r.Exactly(e.JSON(), e.Error())

	e = ValidateP(&z{})
	r.Nil(e)
}

func TestRaw(t *testing.T) {

	r := require.New(t)

	e := Validate(&c{})

	b, err := json.Marshal(e.Data())
	r.Nil(err)

	r.Exactly(json.RawMessage(b), e.Raw())

	e = Validate(&z{})
	r.Nil(e)
}

func TestRawP(t *testing.T) {

	r := require.New(t)

	e := ValidateP(&c{})

	b, err := json.Marshal(e.Data())
	r.Nil(err)

	r.Exactly(json.RawMessage(b), e.Raw())

	e = ValidateP(&z{})
	r.Nil(e)
}
