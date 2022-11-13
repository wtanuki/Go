package Exos

// Returns true if s1 only contains elements from s2

func ContainsOnlyElementsFromString(s1, s2 string) bool {
	rune_s2 := []rune(s2)
	for _, s1char := range s1 {
		if !IsPresentInSliceOfRunes(rune_s2, s1char) {
			return false
		}
	}

	return true
}
