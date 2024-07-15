// TODO: 각 단어의 짝수번째 알파벳은 대문자로, 홀수번째 알파벳은 소문자로 바꾼 문자열을 리턴
// https://school.programmers.co.kr/learn/courses/30/lessons/12930?language=go

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func toWeirdString(s string) string {
	words := strings.Split(s, " ")
	var result []string

	for _, word := range words {
		var newWord []rune

		for i, char := range word {
			if i%2 == 0 {
				newWord = append(newWord, unicode.ToUpper(char))
			} else {
				newWord = append(newWord, unicode.ToLower(char))
			}
		}
		result = append(result, string(newWord))
	}
	return strings.Join(result, " ")
}

func main() {
	fmt.Println(toWeirdString("try hello world")) // Output: "TrY HeLLo WoRLd"
	fmt.Println(toWeirdString("a b c d"))         // Output: "A B C D"
}
