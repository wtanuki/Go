package Exos

import "fmt"

func TabMult(n int) {
	for i := 1; i <= 9; i++ {
		fmt.Printf("%v x %v = %v\n", i, n, i*n)
	}
}
