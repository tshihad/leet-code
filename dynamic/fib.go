package dynamic

var mem = make(map[int]int)

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	if val, ok := mem[n]; ok {
		return val
	}
	mem[n] = fib(n-1) + fib(n-2)

	return mem[n]
}
