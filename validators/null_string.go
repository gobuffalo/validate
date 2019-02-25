package validators

// IsNull defines how packages defines null (empty) strings.
func IsNull(str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}
