package main

import (
	"Exos"
	"fmt"
)

func main() {
	fmt.Println(Exos.IsThereADuplicateInString("0123456789")) // false
	fmt.Println(Exos.IsThereADuplicateInString("0123456780")) // true
	fmt.Println(Exos.IsThereADuplicateInString("choumi"))     // false
	fmt.Println(Exos.IsThereADuplicateInString("chooum"))     // true
}
