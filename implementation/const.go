// TODO: 상수
// https://www.acmicpc.net/problem/2908

package main

import "fmt"

func reverseInt(n int) int {
	reversed := 0
	for n > 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}
	return reversed
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	A = reverseInt(A)
	B = reverseInt(B)

	if A > B {
		fmt.Println(A)
	} else {
		fmt.Println(B)
	}
}
