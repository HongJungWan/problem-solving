// TODO: Two Sum 문제
// TODO: 주어진 배열에서 두 숫자의 합이 특정 목표값(target)과 같아지는 두 숫자의 인덱스를 찾는 문제

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return []int{}
}

/*
[아이디어]

1. numMap[target-num], 맵 numMap에서 키가 target-num인 element를 조회.
2. if j, ok := numMap[target-num]; ok { ... }, Go 언어의 "comma ok" 패턴, 맵 조회 시 특정 키가 존재하는지 확인.
*/
