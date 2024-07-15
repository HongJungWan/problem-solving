// TODO: n을 3진법 상에서 앞뒤로 뒤집은 후, 이를 다시 10진법으로 표현한 수를 return
// https://school.programmers.co.kr/learn/courses/30/lessons/68935

package main

import (
	"fmt"
	"strconv"
)

func base3Conversion(n int) int {
	// 10진법 -> 3진법 변환
	threeBase := strconv.FormatInt(int64(n), 3)

	// 뒤집기
	reversed := reverse(threeBase)

	// 3진법 -> 10진법 변환
	result, _ := strconv.ParseInt(reversed, 3, 64)
	return int(result)
}

func reverse(s string) string {
	runes := []rune(s)
	n := len(s)

	// 문자열을 반으로 나누어 양쪽 끝에서부터 바꾸기
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(base3Conversion(45))  // Output: 7
	fmt.Println(base3Conversion(125)) // Output: 229
}
