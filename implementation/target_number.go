// TODO: 타겟 넘버
// https://school.programmers.co.kr/learn/courses/30/lessons/43165

package main

import "fmt"

func findTargetWayDFS(numbers []int, target int, index int, currentSum int) int {
	// 배열의 끝에 도달했을 때 현재 합계가 목표와 일치하는지 확인
	if index == len(numbers) {
		if currentSum == target {
			return 1
		}
		return 0
	}

	// 현재 숫자를 더하거나 빼는 두 가지 경로를 탐색
	addPath := findTargetWayDFS(numbers, target, index+1, currentSum+numbers[index])
	subtractPath := findTargetWayDFS(numbers, target, index+1, currentSum-numbers[index])

	// 두 경로에서 가능한 방법의 수를 합산
	return addPath + subtractPath
}

func main() {
	numbers1 := []int{1, 1, 1, 1, 1}
	target1 := 3
	result1 := findTargetWayDFS(numbers1, target1, 0, 0)
	fmt.Println(result1) // 출력: 5

	numbers2 := []int{4, 1, 2, 1}
	target2 := 4
	result2 := findTargetWayDFS(numbers2, target2, 0, 0)
	fmt.Println(result2) // 출력: 2
}

/*
문제에서는 주어진 숫자 배열을 사용하여 특정 타겟 숫자를 만들 수 있는 방법의 수를 구해야 한다.

각 숫자에 대해 더하거나 뺄 수 있는 두 가지 선택이 있으며, 이 선택을 통해 가능한 모든 조합을 확인해야 한다.

DFS는 모든 경로를 완전 탐색하여 가능한 모든 조합을 확인할 수 있는 효과적인 방법이다.
*/
