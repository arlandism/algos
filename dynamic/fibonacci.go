package main

// dynamic programming fibonacci

import (
	"fmt"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// [0, 0] fib(0) -> 0
// [0, 1] fib(1) -> 1
// [0, 1] fib(2) -> 1
// [1, 1] fib(3) -> 2
// [1, 2] fib(4) -> 3
// dyfib(4)
// --> 0 + 0 => acc
// --> acc + 1 => acc
// --> acc + 2 => acc
// --> acc + 3 => acc
func dyfib(n int) int {
	x := 0
	y := 1
	var z int
	for i := 0; i < n; i++ {
		x = y     // x = 1, x = 0, x = 1
		y = z     // y = 0, y = 1, y = 1
		z = x + y // z = 1, z = 1, z = 2
	}
	return z
}

func main() {
	fmt.Println(fib(100))
	fmt.Println(dyfib(100))
}
