// TODO: 피보나치
// TODO: 반복문 방식

package main

import "fmt"

func main() {
	n := 10

	fmt.Printf("Fibonacci(%d) = %d\n", n, fibonacciIterative(n))
}

func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

/*
[반복문 방식]
루프를 사용하여 피보나치 수를 계산하는 방법. 시간 복잡도는 O(n)으로 효율적.


a, b = b, a+b

[i=2] a=1, b=1
[i=3] a=1, b=2
[i=4] a=2, b=3
[i=5] a=3, b=5
[i=6] a=5, b=8
[i=7] a=8, b=13
[i=8] a=13, b=21
[i=9] a=21, b=34
[i=10] a=34, b=55
*/
