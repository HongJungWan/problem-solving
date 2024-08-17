// TODO: 기능 개발
// https://school.programmers.co.kr/learn/courses/30/lessons/42586

package main

import "fmt"

func calculateDeployment(progresses []int, speeds []int) []int {
	days := []int{}
	answer := []int{}

	for i, progress := range progresses {
		day := (100 - progress) / speeds[i]
		if (100-progress)%speeds[i] != 0 {
			day++
		}
		days = append(days, day)
	}

	current := days[0]
	count := 1

	for i := 1; i < len(days); i++ {
		if days[i] <= current {
			count++
		} else {
			answer = append(answer, count)
			current = days[i]
			count = 1
		}
	}

	answer = append(answer, count)
	return answer
}

func main() {
	progresses := []int{93, 30, 55}
	speeds := []int{1, 30, 5}
	fmt.Println(calculateDeployment(progresses, speeds)) // [2, 1]
}
