// TODO: 피보나치
// TODO: DP 방식

package main

import "fmt"

func main() {
	n := 10

	fmt.Printf("Fibonacci(%d) = %d\n", n, fibonacciDP(n))
}

func fibonacciDP(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/*
[동적 프로그래밍 방식]
작은 부분 문제를 먼저 해결한 후 이를 이용해 큰 문제를 해결하는 방법. 메모제이션과 유사하지만 반복문을 사용하여 더욱 직관적.
*/
