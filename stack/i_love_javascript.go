// TODO: I Love JavaScript
// https://www.acmicpc.net/problem/27969

package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func parseASON(ason string) int {
	var stack []int
	size := 0

	for i := 0; i < len(ason); i++ {
		char := ason[i]

		switch char {
		case '[':
			stack = append(stack, size)
			size = 0
		case ']':
			size += 8
			if len(stack) > 0 {
				size += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		default:
			if unicode.IsDigit(rune(char)) {
				for i+1 < len(ason) && unicode.IsDigit(rune(ason[i+1])) {
					i++
				}
				size += 8
			} else if unicode.IsLetter(rune(char)) {
				start := i
				for i+1 < len(ason) && unicode.IsLetter(rune(ason[i+1])) {
					i++
				}
				size += (i - start + 1) + 12
			}
		}
	}
	return size
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(parseASON(input))
}
