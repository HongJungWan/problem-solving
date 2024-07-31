// TODO: 전화번호 목록
// https://school.programmers.co.kr/learn/courses/30/lessons/42577

package main

import (
	"fmt"
)

func numberList(phoneBook []string) bool {
	phoneMap := make(map[string]bool)

	// 모든 전화번호를 해시맵에 저장
	for _, phone := range phoneBook {
		phoneMap[phone] = true
	}

	// 각 전화번호에 대해 접두어를 체크
	for _, phone := range phoneBook {
		prefix := ""

		for _, ch := range phone {
			prefix += string(ch)

			// 현재 번호와 동일하지 않은 접두어가 해시맵에 있는지 확인
			if prefix != phone && phoneMap[prefix] {
				return false
			}
		}
	}

	return true
}

func main() {
	phoneBook1 := []string{"119", "97674223", "1195524421"}
	phoneBook2 := []string{"123", "456", "789"}
	phoneBook3 := []string{"12", "123", "1235", "567", "88"}

	fmt.Println(numberList(phoneBook1)) // false
	fmt.Println(numberList(phoneBook2)) // true
	fmt.Println(numberList(phoneBook3)) // false
}
