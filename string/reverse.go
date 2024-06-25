// TODO: 문자열 뒤집기

package main

import "fmt"

func main() {
	input := "hello"
	reversed := reverseString(input)

	fmt.Println("Original:", input)
	fmt.Println("Reversed:", reversed)
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)

	// 문자열을 반으로 나누어 양쪽 끝에서부터 바꾸기
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

/*
[유니코드 지원]

문자열을 rune 슬라이스로 변환하면 유니코드 문자(UTF-8 인코딩된 문자)를 올바르게 처리할 수 있다.
단순히 바이트 단위로 문자열을 뒤집으면 멀티바이트 문자가 깨질 수 있다.
예를 들어, hello 같은 ASCII 문자열 뿐만 아니라, こんにちは 같은 일본어 문자열도 제대로 뒤집을 수 있다.
*/
