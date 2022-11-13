package Exos

// Assume len(s) > 1

func IsThereADuplicateInString(s string) bool {
	var visited []rune
	var rune_s = []rune(s)

	for _, word := range rune_s {
		if !IsPresentInSliceOfRunes(visited, word) {
			visited = append(visited, word)
		}
	}

	return len(visited) != len(rune_s)
}
