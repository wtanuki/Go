package Exos

func IsAFullDigitString(s string) bool {

	for _, word := range s {
		if word < '0' || word > '9' {
			return false
		}
	}

	return true
}
