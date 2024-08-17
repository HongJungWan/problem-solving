// TODO: 짝수를 찾아라
// https://www.acmicpc.net/problem/3058

package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scan(&T) // 입력할 테스트 데이터 수

	for i := 0; i < T; i++ {
		var sum, min, num int
		min = math.MaxInt32 // 초기 최솟값 설정, 가능한 가장 큰 값

		for j := 0; j < 7; j++ {
			fmt.Scan(&num) // 각 자연수 입력

			if num%2 == 0 { // 짝수인지 확인
				sum += num // 짝수의 합

				if num < min {
					min = num // 짝수 중 최솟값 갱신
				}
			}
		}
		fmt.Printf("%d %d\n", sum, min) // 결과 출력
	}
}
