// TODO: 각 알파벳을 일정한 거리만큼 밀어서 다른 알파벳으로 바꾸는 암호화 방식
// https://school.programmers.co.kr/learn/courses/30/lessons/12926

package main

import (
	"fmt"
	"unicode"
)

func toCaesarCipher(s string, n int) string {
	var result string

	for _, char := range s {
		if unicode.IsLower(char) {
			result += string((char-'a'+rune(n))%26 + 'a')
		} else if unicode.IsUpper(char) {
			result += string((char-'A'+rune(n))%26 + 'A')
		} else {
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(toCaesarCipher("AB", 1))    // "BC"
	fmt.Println(toCaesarCipher("z", 1))     // "a"
	fmt.Println(toCaesarCipher("a B z", 4)) // "e F d"
}

/*
시저 암호화 공식
1. new_char = ((char - 'a' + n) % 26) + 'a'
2. new_char = ((char - 'A' + n) % 26) + 'A'
*/
