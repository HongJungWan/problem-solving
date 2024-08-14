// TODO: 체육복
// https://school.programmers.co.kr/learn/courses/30/lessons/42862

package main

import (
	"fmt"
)

func gymClothesRent(n int, lost []int, reserve []int) int {
	// 모든 학생이 체육복을 가지고 있다고 가정
	students := make([]int, n)

	for _, v := range lost {
		students[v-1]-- // 체육복을 도난당한 학생들
	}
	for _, v := range reserve {
		students[v-1]++ // 여벌의 체육복을 가진 학생들
	}

	// 체육복 빌려주기
	for i := 0; i < n; i++ {
		if students[i] == -1 { // 체육복이 없는 학생일 경우

			if i > 0 && students[i-1] == 1 { // 앞 번호 학생이 여벌이 있는지 확인
				students[i]++
				students[i-1]--
			} else if i < n-1 && students[i+1] == 1 { // 뒷 번호 학생이 여벌이 있는지 확인
				students[i]++
				students[i+1]--
			}
		}
	}

	// 체육 수업에 참여할 수 있는 학생 수 계산
	count := 0
	for i := 0; i < n; i++ {
		if students[i] >= 0 {
			count++
		}
	}

	return count
}

func main() {
	n := 5
	lost := []int{2, 4}
	reserve := []int{1, 3, 5}
	fmt.Println(gymClothesRent(n, lost, reserve)) // 5
}
