package main

import (
	"Exos"
	"fmt"
)

func main() {
	fmt.Println(Exos.Atoi("12345"))
	fmt.Println(Exos.Atoi("0000000012345"))
	fmt.Println(Exos.Atoi("012 345"))
	fmt.Println(Exos.Atoi("Hello World!"))
	fmt.Println(Exos.Atoi("+1234"))
	fmt.Println(Exos.Atoi("-1234"))
	fmt.Println(Exos.Atoi("++1234"))
	fmt.Println(Exos.Atoi("--1234"))
}
