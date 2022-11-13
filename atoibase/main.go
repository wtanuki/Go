package main

import (
	"Exos"
	"fmt"
)

func main() {
	fmt.Println(Exos.AtoiBase("125", "0123456789"))      // 125
	fmt.Println(Exos.AtoiBase("1111101", "01"))          // 125
	fmt.Println(Exos.AtoiBase("7D", "0123456789ABCDEF")) // 125
	fmt.Println(Exos.AtoiBase("uoi", "choumi"))          // 125
	fmt.Println(Exos.AtoiBase("bbbbbab", "-ab"))         // 0
}
