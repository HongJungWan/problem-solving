// TODO: 점프 게임

/*
문제 설명:
주어진 정수 배열 nums에서, 각 인덱스의 숫자는 그 위치에서 최대로 점프할 수 있는 거리를 나타냅니다.
배열의 시작 지점에서 출발하여 배열의 마지막 위치에 도달할 수 있는지를 판별하세요.

예시 입력: nums = [2,3,1,1,4]

예시 출력: true
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println("Can jump to the end:", canJump(nums))
}

func canJump(nums []int) bool {
	maxReach := 0

	for i, num := range nums {
		if i > maxReach {
			return false
		}
		if i+num > maxReach {
			maxReach = i + num
		}
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}
