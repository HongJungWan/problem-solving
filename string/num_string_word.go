// TODO: 숫자 문자열과 영단어
// https://school.programmers.co.kr/learn/courses/30/lessons/81301?language=go

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func toIntConversion(s string) int {
	numWords := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for key, value := range numWords {
		s = strings.ReplaceAll(s, key, value)
	}

	result, _ := strconv.Atoi(s)
	return result
}

func main() {
	fmt.Println(toIntConversion("one4seveneight"))
	fmt.Println(toIntConversion("23four5six7"))
	fmt.Println(toIntConversion("1zerotwozero3"))
}
