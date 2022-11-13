package main

import (
	"Exos"
	"fmt"
)

func main() {
	fmt.Println(Exos.ContainsOnlyElementsFromString("125", "0123456789"))       // true
	fmt.Println(Exos.ContainsOnlyElementsFromString("1111101", "01"))           // true
	fmt.Println(Exos.ContainsOnlyElementsFromString("7DX", "0123456789ABCDEF")) // false
	fmt.Println(Exos.ContainsOnlyElementsFromString("Zuoi", "choumi"))          // false
}
