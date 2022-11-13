package Exos

// Assume n and p are positive int

func Pow(n, p int) int {
	var res int = n

	for i := 1; i < p; i++ {
		res *= n
	}
	return res
}
