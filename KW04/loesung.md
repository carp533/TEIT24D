
func ackermann(m, n uint) uint {
	if m == 0 {
		return n + 1
	}
	if n == 0 {
		return ackermann(m-1, 1)
	}
	return ackermann(m-1, ackermann(m, n-1))
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibList(n int) []int {
	fibs := make([]int, n+1)
	fibs[0], fibs[1] = 0, 1

	for i := 2; i <= n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs[1:]
}

func fibBig(n uint64) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}
	a, b := big.NewInt(0), big.NewInt(1)
	for n--; n > 0; n-- {
		a.Add(a, b)
		a, b = b, a
	}
	return b
}

// Prüft, ob eine Zahl gerade ist
func isEven(n int) bool {
	if n == 0 {
		return true
	}
	return isOdd(n - 1)
}

// Prüft, ob eine Zahl ungerade ist
func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

func mult(a, b uint) uint {
	if a == 0 {
		return 0
	}
	return mult(a-1, b) + b
}

