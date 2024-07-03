// TODO: 괄호 문제

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()

	if isValidParentheses(inputString) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isValidParentheses(s string) bool {
	stack := []rune{}

	for _, char := range s {

		if char == '(' {
			stack = append(stack, char)

		} else if char == ')' {

			if len(stack) == 0 {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
