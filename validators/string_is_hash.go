package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsHash is a validator object.
type StringIsHash struct {
	Name      string
	Field     string
	Algorithm string
}

// Validate adds an error if the Field is not formatted as a hash of provided type algorithm.
// Algorithm can be one of ['md4', 'md5', 'sha1', 'sha256', 'sha384', 'sha512', 'ripemd128',
// 'ripemd160', 'tiger128', 'tiger160', 'tiger192', 'crc32', 'crc32b'].
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

	e.Add(v.Name, fmt.Sprintf("%s was not evaluated as valid %s hash", v.Name, v.Algorithm))
}

// SetField sets validator field.
func (v *StringIsHash) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsHash) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
