// TODO: 백대열
// https://www.acmicpc.net/problem/14490

package main

import (
	"fmt"
	"os"
)

// 두 정수의 최대공약수(GCD)를 구하는 메서드 (유클리드 호제법)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var n, m int

	// 입력 받기: 'n:m' 형태로 두 정수를 입력받음
	_, err := fmt.Scanf("%d:%d", &n, &m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "입력을 읽는 데 실패했습니다: %v\n", err)
		os.Exit(1)
	}

	// 최대공약수(GCD) 계산
	gcdValue := gcd(n, m)

	// 최대공약수로 나눠 간단한 형태의 분수로 축약
	n /= gcdValue
	m /= gcdValue

	fmt.Printf("%d:%d\n", n, m)
}

// 유클리드 호제법 : 두 자연수의 최대공약수(GCD)를 구하는 알고리즘.
// 이 방법은 두 수 중 큰 수를 작은 수로 나눈 나머지를 다시 작은 수와 나누는 과정을 반복하여, 나머지가 0이 될 때의 나누는 수가 최대공약수.
