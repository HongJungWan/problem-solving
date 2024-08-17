// TODO: 쇠막대기
// https://www.acmicpc.net/problem/10799

package main

import (
	"bufio"
	"fmt"
	"os"
)

func processBracket(i int, expression string, stack []int, result int) int {
	// 레이저에 의해 잘리는 쇠막대기의 수
	if i > 0 && expression[i-1] == '(' {
		stack = stack[:len(stack)-1]
		result += len(stack)

	} else { // 쇠막대기의 끝에 도달했을 때, 마지막 조각을 결과에 추가하는 경우
		stack = stack[:len(stack)-1]
		result += 1
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')

	var stack []int
	result := 0

	for i := 0; i < len(expression)-1; i++ {
		if expression[i] == '(' {
			stack = append(stack, i)
		} else if expression[i] == ')' {
			result = processBracket(i, expression, stack, result)
		}
	}
	fmt.Println(result)
}
