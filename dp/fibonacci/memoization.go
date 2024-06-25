// TODO: 피보나치
// TODO: 메모제네이션 방식

package main

import "fmt"

func main() {
	n := 10

	memoization := make(map[int]int)
	fmt.Printf("Fibonacci(%d) = %d\n", n, fibonacciMemoization(n, memoization))
}

func fibonacciMemoization(n int, memoization map[int]int) int {
	if n <= 1 {
		return n
	}

	if value, exists := memoization[n]; exists {
		return value
	}

	memoization[n] = fibonacciMemoization(n-1, memoization) + fibonacciMemoization(n-2, memoization)
	return memoization[n]
}

/*
[메모제이션 방식]
재귀 호출 시 이미 계산된 값을 저장하여 중복 계산을 피한다. 공간과 시간 복잡도 모두 효율적.

1. memoization 변수는 공유 자원이다.

2. comma ok 패턴을 사용하여 키가 존재하는지 확인, 키가 존재하면 해당 값을 반환하고, 존재하지 않으면 피보나치 수를 계산하여 memoization 맵에 저장한 후 반환.
*/
