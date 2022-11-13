package Exos

// Assume n is a positive int

func Pow(n, p int) int {
	var res int = n

	if p == 0 {
		return 1
	}

	for i := 1; i < p; i++ {
		res *= n
	}
	return res
}
