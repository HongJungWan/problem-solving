// TODO: 모의고사
// TODO:
// https://school.programmers.co.kr/learn/courses/30/lessons/42840

package main

import (
	"fmt"
)

var supoza1 = []int{1, 2, 3, 4, 5}
var supoza2 = []int{2, 1, 2, 3, 2, 4, 2, 5}
var supoza3 = []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}

func solution(answers []int) []int {
	scores := make([]int, 3)

	for i, answer := range answers {
		if answer == supoza1[i%len(supoza1)] {
			scores[0]++
		}
		if answer == supoza2[i%len(supoza2)] {
			scores[1]++
		}
		if answer == supoza3[i%len(supoza3)] {
			scores[2]++
		}
	}

	maxScore := max(scores[0], max(scores[1], scores[2]))

	var result []int
	for i, score := range scores {
		if score == maxScore {
			result = append(result, i+1)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(solution([]int{1, 2, 3, 4, 5})) // Output: [1]
	fmt.Println(solution([]int{1, 3, 2, 4, 2})) // Output: [1,2,3]
}
