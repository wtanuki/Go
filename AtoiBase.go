package Exos

/* Write a function that takes two arguments:

s: a numeric string in a given base.
base: a string representing all the different digits that can represent a numeric value.
And return the integer value of s in the given base.

If the base is not valid it returns 0.

Validity rules for a base :

A base must contain at least 2 characters.
Each character of a base must be unique.
A base should not contain + or - characters.
String number must contain only elements that are in base.

Only valid string numbers will be tested.

The function does not have to manage negative numbers */

func AtoiBase(s string, base string) int {
	var res int = 0
	var idxBase int = 0
	p := 0
	basePowP := 0

	if IsAValidBase(base) && ContainsOnlyElementsFromString(s, base) {
		new_base := len(base)
		for idxS, word := range s {
			idxBase = FindIndex(base, word)
			p = (len(s) - (idxS + 1))
			basePowP = Pow(new_base, p)
			res = idxBase*basePowP + res
		}
	}
	return res
}
