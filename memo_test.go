package memo

import (
	"fmt"
	"time"
)

func ExampleFib() {
	var fib func(int) int
	fib = func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		return fib(n-1) + fib(n-2)
	}
	memoFib := Memoize(fib)

	start := time.Now()
	out := memoFib(45)
	fmt.Println(out, time.Since(start))

	startCached := time.Now()
	out = memoFib(45)
	fmt.Println(out, time.Since(startCached))
	// Output:
}
