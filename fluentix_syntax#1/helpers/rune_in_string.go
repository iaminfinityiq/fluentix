package helpers

func RuneInString(r rune, s string) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}

	return false
}