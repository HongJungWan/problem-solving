// TODO: 중복된 문자 제거
// https://school.programmers.co.kr/learn/courses/30/lessons/120888

package main

import (
	"fmt"
)

func removeDuplicate(my_string string) string {
	checkDuplicates := make(map[rune]bool)
	var temp []rune

	for _, char := range my_string {
		if !checkDuplicates[char] {
			temp = append(temp, char)
			checkDuplicates[char] = true
		}
	}

	return string(temp)
}

func main() {
	my_string1 := "people"
	my_string2 := "We are the world"

	fmt.Println(removeDuplicate(my_string1)) // "peol"
	fmt.Println(removeDuplicate(my_string2)) // "We arthwold"
}
