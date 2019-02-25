package validators

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsHash(t *testing.T) {

	r := require.New(t)
	e := validator.NewErrors()

	h := md5.New()
	_, err := io.WriteString(h, "Hello World!")
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	// md5
	v := StringIsHash{Name: "Name", Algorithm: "md5", Field: fmt.Sprintf("%x", h.Sum(nil))}
	v.Validate(e)
	r.Equal(0, e.Count())

	h = sha256.New()
	_, err = io.WriteString(h, "Hello World!")
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	// sha256
	v = StringIsHash{Name: "Name", Algorithm: "sha256", Field: fmt.Sprintf("%x", h.Sum(nil))}
	v.Validate(e)
	r.Equal(0, e.Count())

	// empty algorithm is invalid
	v = StringIsHash{Name: "Name", Algorithm: "", Field: fmt.Sprintf("%x", h.Sum(nil))}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{" is an unknown hash algorithm"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// unknown algorithm is invalid
	v = StringIsHash{Name: "Name", Algorithm: "unknown", Field: fmt.Sprintf("%x", h.Sum(nil))}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"unknown is an unknown hash algorithm"}, e.Get("Name"))

	// reset
	e = validator.NewErrors()

	// random string is invalid
	v = StringIsHash{Name: "Name", Algorithm: "md5", Field: "whatisthis"}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name was not evaluated as valid md5 hash"}, e.Get("Name"))
}
