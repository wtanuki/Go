package main

import (
	"Exos"
	"fmt"
)

func main() {
	fmt.Println(Exos.BasicAtoi2("12345"))
	fmt.Println(Exos.BasicAtoi2("0000000012345"))
	fmt.Println(Exos.BasicAtoi2("012 345"))
	fmt.Println(Exos.BasicAtoi2("Hello World!"))
}
