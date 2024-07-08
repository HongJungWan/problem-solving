package main

import (
	"fmt"
	"math"
)

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

func main() {
	L := 6
	R := 10

	result := countPrimeSetBits(L, R)
	fmt.Printf("The number of integers in the range [%d, %d] with a prime number of set bits is: %d\n", L, R, result)
}
