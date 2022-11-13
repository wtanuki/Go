package Exos

func FindIndex(s string, r rune) int {
	for idx, word := range s {
		if word == r {
			return idx
		}
	}
	return -1
}
