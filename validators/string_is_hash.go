package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsHash is a validator object
type StringIsHash struct {
	Name      string
	Field     string
	Message   string
	Algorithm string
}

// Validate checks if a string is a hash of type algorithm.
// Algorithm is one of ['md4', 'md5', 'sha1', 'sha256', 'sha384', 'sha512', 'ripemd128', 'ripemd160', 'tiger128', 'tiger160', 'tiger192', 'crc32', 'crc32b']
func (v *StringIsHash) Validate(e *validator.Errors) {

	length := "0"
	algo := strings.ToLower(v.Algorithm)

	switch algo {
	case "crc32", "crc32b":
		length = "8"
	case "md5", "md4", "ripemd128", "tiger123":
		length = "32"
	case "sha1", "ripemd160", "tiger160":
		length = "40"
	case "tiger192":
		length = "48"
	case "sha256":
		length = "64"
	case "sha384":
		length = "96"
	case "sha512":
		length = "128"
	default:
		e.Add(v.Name, fmt.Sprintf("%s is an unknown hash algorithm", v.Algorithm))
		return
	}

	matched, err := regexp.MatchString("^[a-f0-9]{"+length+"}$", v.Field)

	if matched && err == nil {
		return
	}

	// adding custom error message
	if len(v.Message) > 0 {
		e.Add(v.Name, v.Message)
		return
	}

	// or standard message
	e.Add(v.Name, fmt.Sprintf("%s was not evaluated as valid %s hash", v.Name, v.Algorithm))
}
