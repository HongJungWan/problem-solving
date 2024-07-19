// TODO: 모의고사
// TODO: 수포자들이 정답을 맞추는 패턴을 정의하고, 각 수포자의 패턴을 반복하여 주어진 답안과 비교한 후 가장 많은 정답을 맞춘 수포자를 찾는다.
// https://school.programmers.co.kr/learn/courses/30/lessons/42840

package main

import (
	"fmt"
)

// 수포자들의 답안 패턴을 명확한 변수명으로 정의
var studentOnePattern = []int{1, 2, 3, 4, 5}
var studentTwoPattern = []int{2, 1, 2, 3, 2, 4, 2, 5}
var studentThreePattern = []int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}

func findTopScorers(answers []int) []int {
	scores := calculateScores(answers)
	return identifyTopScorers(scores)
}

// 가장 많은 정답을 맞춘 수포자를 찾음
func identifyTopScorers(scores []int) []int {
	maxScore := max(scores...)
	var topScorers []int

	for i, score := range scores {
		if score == maxScore {
			topScorers = append(topScorers, i+1)
		}
	}
	return topScorers
}

// 각 수포자의 점수를 계산
func calculateScores(answers []int) []int {
	scores := []int{0, 0, 0}

	for i, answer := range answers {
		if answer == studentOnePattern[i%len(studentOnePattern)] {
			scores[0]++
		}
		if answer == studentTwoPattern[i%len(studentTwoPattern)] {
			scores[1]++
		}
		if answer == studentThreePattern[i%len(studentThreePattern)] {
			scores[2]++
		}
	}

	return scores
}

// 가변 인자를 받아 최대값을 반환
func max(nums ...int) int {
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func main() {
	fmt.Println(findTopScorers([]int{1, 2, 3, 4, 5})) // Output: [1]
	fmt.Println(findTopScorers([]int{1, 3, 2, 4, 2})) // Output: [1,2,3]
}
