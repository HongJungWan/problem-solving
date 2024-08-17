// TODO: 좋은 단어
// https://www.acmicpc.net/problem/3986

package main

import (
	"bufio"
	"fmt"
	"os"
)

func isGoodWord(word string) bool {
	stack := make([]rune, 0, len(word)/2)

	for _, char := range word {
		if len(stack) > 0 && stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func main() {
	var n, goodWords int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	for ; n > 0; n-- {
		word, _ := reader.ReadString('\n')
		word = word[:len(word)-1]

		if isGoodWord(word) {
			goodWords++
		}
	}
	fmt.Println(goodWords)
}
