package Exos

/* atoi
Instructions
Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number represented as a string in a number represented as an int.

Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.

For this exercise the handling of the signs + or - does have to be taken into account.

This function will only have to return the int. For this exercise the error result of Atoi is not required. */

func Atoi(s string) int {
	var res int = 0
	rune_s := []rune(s)
	signe := 1

	if len(s) > 0 {

		switch s[0] {
		case '+':
			rune_s = rune_s[1:]
			break
		case '-':
			rune_s = rune_s[1:]
			signe = -1
			break
		}

	}

	res = BasicAtoi2(string(rune_s)) * signe
	return res
}
