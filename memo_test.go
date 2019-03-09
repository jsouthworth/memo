package memo

import (
	"fmt"
	"time"
)

func ExampleFib() {
	// This will cache the final results of the fib
	// function significantly speeding up subsequent operations
	// with the expense of some memory. Obviously memoization
	// can only be done on effect-free functions.
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
}

func ExampleFibMemo() {
	// This will cache all intermediate results of the
	// fib function and lead to much faster computation
	// of the initial result as well at the expense of
	// more memory.
	var memoFib func(...interface{}) interface{}
	memoFib = Memoize(func(n int) int {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
		return memoFib(n-1).(int) +
			memoFib(n-2).(int)
	})

	start := time.Now()
	out := memoFib(45)
	fmt.Println(out, time.Since(start))

	startCached := time.Now()
	out = memoFib(45)
	fmt.Println(out, time.Since(startCached))
}
