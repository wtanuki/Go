package main

import (
	"Exos"
	"fmt"
)

func main() {
	fmt.Println(Exos.IsAValidBase("0123456789"))       // true
	fmt.Println(Exos.IsAValidBase("01"))               // true
	fmt.Println(Exos.IsAValidBase("0123456789ABCDEF")) // true
	fmt.Println(Exos.IsAValidBase("choumi"))           // true

	fmt.Println(Exos.IsAValidBase("-ab")) // false
}
