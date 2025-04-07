package helpers

func StringToRune(s string) []rune {
	var returned []rune = []rune{}

	for _, r := range s {
		returned = append(returned, r)
	}

	return returned
}