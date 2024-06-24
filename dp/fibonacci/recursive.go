// TODO: 피보나치
// TODO: 재귀 방식

package main

import "fmt"

func main() {
	n := 10

	fmt.Printf("Fibonacci(%d) = %d\n", n, fibonacciRecursive(n))
}

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

/*
[재귀 방식]
가장 간단하지만 비효율적인 방법.
같은 계산을 여러 번 반복하기 때문에 시간 복잡도가 매우 높다.
*/
