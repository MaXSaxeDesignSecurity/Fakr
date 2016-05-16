package locale

// Verify is used to validate that a locale is supported by Fakr
func Verify(locale string) bool {
	valid := false
	for _, possible := range List() {
		if possible == locale {
			valid = true
		}
	}

	return valid
}
