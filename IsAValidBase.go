package Exos

/* Validity rules for a base :

A base must contain at least 2 characters.
Each character of a base must be unique.
A base should not contain + or - characters. */

func IsAValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for _, word := range base {
		if word == '+' || word == '-' {
			return false
		}
	}
	if IsThereADuplicateInString(base) == true {
		return false
	}

	return true
}
