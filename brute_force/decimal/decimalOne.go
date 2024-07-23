// TODO: 이진 비트 수가 소수인 숫자 개수 구하기, 비트 연산
// https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/description/

package main

import (
	"fmt"
	"math"
)

// 소수 판별
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 이진수 변환 및 비트 카운팅
func countSetBits(n int) int {
	count := 0

	for n > 0 {
		count += n & 1 // 마지막 비트가 1인지 확인
		n >>= 1        // 오른쪽으로 1 비트 이동
	}
	return count
}

// L부터 R까지 반복
func countPrimeSetBits(L int, R int) int {
	primeCount := 0

	for i := L; i <= R; i++ {
		if isPrime(countSetBits(i)) {
			primeCount++
		}
	}
	return primeCount
}

func main() {
	L := 6
	R := 10

	result := countPrimeSetBits(L, R)
	fmt.Printf("범위 [%d, %d] 내에서 이진수로 표현했을 때 1로 설정된 비트 수가 소수인 정수의 개수는: %d\n", L, R, result)
}
