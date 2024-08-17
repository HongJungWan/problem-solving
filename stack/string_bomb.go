// TODO: 문자열 폭발
// https://www.acmicpc.net/problem/9935

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s, bomb string
	fmt.Fscan(reader, &s)
	fmt.Fscan(reader, &bomb)

	bombLen := len(bomb)
	stack := []rune{}
	for _, ch := range s {
		stack = append(stack, ch)

		// 스택의 길이가 폭발 문자열의 길이보다 크거나 같고, 스택의 끝 부분이 폭발 문자열과 일치하는 경우
		if len(stack) >= bombLen && string(stack[len(stack)-bombLen:]) == bomb {
			// 폭발 문자열이 스택의 끝에 존재하는 경우 제거
			stack = stack[:len(stack)-bombLen]
		}
	}

	if len(stack) == 0 {
		fmt.Fprintln(writer, "FRULA")
	} else {
		fmt.Fprintln(writer, string(stack))
	}
}
