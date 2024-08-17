// TODO: 부녀회장이 될테야
// https://www.acmicpc.net/problem/2775

package main

import (
	"fmt"
)

func main() {
	var T, k, n int
	fmt.Scan(&T)
	results := make([]int, T)

	// 계산 범위 내의 최대 층과 호수를 설정
	maxK := 14
	maxN := 14

	// 사람 수를 저장할 배열 생성 및 초기화
	residents := make([][]int, maxK+1)
	for i := range residents {
		residents[i] = make([]int, maxN+1)
	}

	// 0층 초기화: i호에는 i명
	for i := 1; i <= maxN; i++ {
		residents[0][i] = i
	}

	// 1층부터 max K층까지 계산
	for floor := 1; floor <= maxK; floor++ {
		for room := 1; room <= maxN; room++ {
			sum := 0
			// 해당 층의 해당 호수에 사는 사람 수 계산
			for j := 1; j <= room; j++ {
				sum += residents[floor-1][j]
			}
			residents[floor][room] = sum
		}
	}

	// 입력받은 케이스에 대해 결과 계산
	for i := 0; i < T; i++ {
		fmt.Scan(&k)
		fmt.Scan(&n)
		results[i] = residents[k][n]
	}

	// 결과 출력
	for _, result := range results {
		fmt.Println(result)
	}
}
