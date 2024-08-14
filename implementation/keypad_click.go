// TODO: 키패드 누르기
// https://school.programmers.co.kr/learn/courses/30/lessons/67256

package main

import (
	"fmt"
	"math"
)

// 숫자 좌표 정의
var position = map[int][2]int{
	1: {0, 0}, 2: {0, 1}, 3: {0, 2},
	4: {1, 0}, 5: {1, 1}, 6: {1, 2},
	7: {2, 0}, 8: {2, 1}, 9: {2, 2},
	0: {3, 1}, '*': {3, 0}, '#': {3, 2},
}

// 두 좌표 간의 맨해튼 거리를 계산
func distance(pos1, pos2 [2]int) int {
	return int(math.Abs(float64(pos1[0]-pos2[0])) + math.Abs(float64(pos1[1]-pos2[1])))
}

func keypadDistance(numbers []int, hand string) string {
	leftPosition := position['*']  // 왼손의 초기 위치 (*)
	rightPosition := position['#'] // 오른손의 초기 위치 (#)
	result := ""

	for _, number := range numbers {
		// 숫자가 왼쪽 열에 위치한 경우
		if number == 1 || number == 4 || number == 7 {
			result += "L"
			leftPosition = position[number]
		} else if number == 3 || number == 6 || number == 9 {
			// 숫자가 오른쪽 열에 위치한 경우
			result += "R"
			rightPosition = position[number]

		} else {
			// 숫자가 가운데 열에 위치한 경우
			leftDistance := distance(leftPosition, position[number])
			rightDistance := distance(rightPosition, position[number])

			if leftDistance < rightDistance {
				result += "L"
				leftPosition = position[number]
			} else if rightDistance < leftDistance {
				result += "R"
				rightPosition = position[number]

			} else {
				// 거리가 같은 경우 주 손잡이에 따라 결정
				if hand == "left" {
					result += "L"
					leftPosition = position[number]
				} else {
					result += "R"
					rightPosition = position[number]
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(keypadDistance([]int{1, 3, 4, 5, 8, 2, 1, 4, 5, 9, 5}, "right")) // 출력: "LRLLLRLLRRL"
	fmt.Println(keypadDistance([]int{7, 0, 8, 2, 8, 3, 1, 5, 7, 6, 2}, "left"))  // 출력: "LRLLRRLLLRR"
	fmt.Println(keypadDistance([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, "right"))    // 출력: "LLRLLRLLRL"
}
