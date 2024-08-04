// TODO: 스킬트리
// https://school.programmers.co.kr/learn/courses/30/lessons/49993

package main

import (
	"fmt"
	"strings"
)

func skillTreeCount(skill string, skillTrees []string) int {
	count := 0

	for _, skillTree := range skillTrees {
		if isValidSkillTree(skill, skillTree) {
			count++
		}
	}

	return count
}

func isValidSkillTree(skill string, skillTree string) bool {
	// 스킬 트리에서 선행 스킬 순서에 있는 스킬들만 추출
	filtered := ""

	for _, s := range skillTree {
		if strings.Contains(skill, string(s)) {
			filtered += string(s)
		}
	}

	// 필터링된 스킬 트리와 실제 선행 스킬 순서를 비교
	return strings.HasPrefix(skill, filtered)
}

func main() {
	skill := "CBD"
	skillTrees := []string{"BACDE", "CBADF", "AECB", "BDA"}

	result := skillTreeCount(skill, skillTrees)
	fmt.Println(result) // 출력: 2
}
