package diffsquares

func SquareOfSum(n int) (r int) {
	for i := 1; i <= n; i++ {
		r += i
	}
	return r * r
}

func SumOfSquares(n int) (r int) {
	for i := 1; i <= n; i++ {
		r += i * i
	}
	return r
}

func Difference(n int) (r int) {
	return SquareOfSum(n) - SumOfSquares(n)
}
