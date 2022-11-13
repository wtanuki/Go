package main

import (
	"Exos"
	"fmt"
	"os"
	"strconv"
)

func main() {
	Args := os.Args[1:]
	if len(Args) == 1 {
		arg, err := strconv.Atoi(Args[0])
		if err != nil {
			fmt.Println(err)
		} else {
			Exos.TabMult(arg)
		}
	}
}
