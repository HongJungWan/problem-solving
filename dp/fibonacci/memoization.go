// TODO: 피보나치
// TODO: 메모제네이션 방식

package main

import "fmt"

func main() {
	n := 10

	memo := make(map[int]int)
	fmt.Printf("Fibonacci(%d) = %d\n", n, fibonacciMemoization(n, memo))
}

func fibonacciMemoization(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if _, exists := memo[n]; !exists {
		memo[n] = fibonacciMemoization(n-1, memo) + fibonacciMemoization(n-2, memo)
	}
	return memo[n]
}

/*
[메모제이션 방식]
재귀 호출 시 이미 계산된 값을 저장하여 중복 계산을 피한다. 공간과 시간 복잡도 모두 효율적.
*/
