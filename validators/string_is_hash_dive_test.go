package validators

import (
	"crypto/sha256"
	"fmt"
	"io"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsHashDive(t *testing.T) {

	r := require.New(t)

	h := sha256.New()
	_, err := io.WriteString(h, "Hello World!")
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	hs := fmt.Sprintf("%x", h.Sum(nil))

	field := []string{hs, hs, hs}

	e := validator.NewErrors()
	v := StringSliceDive{
		Validator: &StringIsHash{
			Algorithm: "sha256",
			Name:      "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{hs, hs + " ", "", " ", "as30"}

	e = validator.NewErrors()
	v = StringSliceDive{
		Validator: &StringIsHash{
			Algorithm: "sha256",
			Name:      "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())
}
