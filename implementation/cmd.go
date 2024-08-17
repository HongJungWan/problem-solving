// TODO: 명령 프롬프트
// https://www.acmicpc.net/problem/1032

package main

import (
	"bufio"
	"fmt"
	"os"
)

func findPattern(files []string) string {
	if len(files) == 0 {
		return ""
	}

	pattern := make([]rune, len(files[0]))
	for i := range pattern {
		pattern[i] = rune(files[0][i])
	}

	for _, file := range files[1:] {
		for i, ch := range file {
			if pattern[i] != ch {
				pattern[i] = '?'
			}
		}
	}
	return string(pattern)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 첫 줄(파일 개수) 읽기, 여기서는 사용하지 않음

	var files []string
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	pattern := findPattern(files)
	fmt.Println(pattern)
}
