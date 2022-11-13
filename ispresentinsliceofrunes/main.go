package main

import (
	"Exos"
	"fmt"
)

func main() {
	var j = "abcdefghijklmnopqrstuvwxyz"
	s := []rune(j)

	fmt.Println(Exos.IsPresentInSliceOfRunes(s, 'g')) // true
	fmt.Println(Exos.IsPresentInSliceOfRunes(s, 'A')) // false
	fmt.Println(Exos.IsPresentInSliceOfRunes(s, '?')) // false
	fmt.Println(Exos.IsPresentInSliceOfRunes(s, '0')) // false
	fmt.Println(Exos.IsPresentInSliceOfRunes(s, 'z')) // true
}
