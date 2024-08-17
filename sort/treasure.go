// TODO: 보물
// https://www.acmicpc.net/problem/1026

package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	B := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&B[i])
	}

	// A를 오름차순으로 정렬
	sort.Ints(A)

	// B를 내림차순으로 정렬
	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	S := 0
	for i := 0; i < N; i++ {
		S += A[i] * B[i]
	}
	fmt.Println(S)
}

// sort.Ints() 함수는 int 타입의 슬라이스를 오름차순으로 정렬

// sort.Slice(slice, func(i, j int) bool {
// 	// 비교 로직, i번째 원소가 j번째 원소보다 "작으면" true 반환
// })
