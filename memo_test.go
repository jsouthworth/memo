package memo

import (
	"fmt"
	"testing"
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

	out := memoFib(45)
	fmt.Println(out)

	out = memoFib(45)
	fmt.Println(out)
	// Output:1134903170
	// 1134903170
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

	out := memoFib(45)
	fmt.Println(out)

	out = memoFib(45)
	fmt.Println(out)
	// Output:1134903170
	// 1134903170
}

func BenchmarkFib(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		memoFib(i % 45)
	}

}

func BenchmarkMemoFib(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		memoFib(i % 45)
	}

}

func BenchmarkFastFibRecursive(b *testing.B) {
	var fibIter func(a, b, p, q, count int) int
	fibIter = func(a, b, p, q, count int) int {
		switch {
		case count == 0:
			return b
		case count%2 == 0:
			return fibIter(
				a,
				b,
				(p*p)+(q*q),
				(2*p*q)+(q*q),
				count/2)
		default:
			return fibIter(
				(b*q)+(a*q)+(a*p),
				(b*p)+(a*q),
				p,
				q,
				count-1)
		}
	}
	fib := func(n int) int {
		return fibIter(1, 0, 0, 1, n)
	}
	for i := 0; i < b.N; i++ {
		fib(i % 45)
	}
}

func BenchmarkFastFib(b *testing.B) {
	fib := func(n int) int {
		a, b, p, q, count := 1, 0, 0, 1, n
		for count != 0 {
			switch {
			case count%2 == 0:
				p, q, count = (p*p)+(q*q), (2*p*q)+(q*q), count/2
			default:
				a, b, count = (b*q)+(a*q)+(a*p), (b*p)+(a*q), count-1
			}
		}
		return b
	}
	for i := 0; i < b.N; i++ {
		fib(i % 45)
	}
}
