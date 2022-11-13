package Exos

// Instructions
// Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number defined as a string in a number defined as an int.

// Atoi returns 0 if the string is not considered as a valid number. For this exercise only valid string will be tested. They will only contain one or several digits as characters.

// For this exercise the handling of the signs + or - does not have to be taken into account.

// This function will only have to return the int. For this exercise the error return of Atoi is not required.

func BasicAtoi(s string) int {
	var res int = 0
	var digit int = 0

	for _, word := range s {
		digit = int(word) - 48
		res = res*10 + digit
		digit = 0
	}
	return res
}
