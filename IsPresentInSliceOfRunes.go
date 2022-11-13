package Exos

func IsPresentInSliceOfRunes(s []rune, r rune) bool {

	for _, word := range s {
		if word == r {
			return true
		}
	}

	return false
}
