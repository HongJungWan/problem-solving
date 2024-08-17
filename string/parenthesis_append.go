// TODO: 괄호 끼워넣기
// https://www.acmicpc.net/problem/11899

package main

import (
	"bufio"
	"fmt"
	"os"
)

func processBracketSequence(inputStr string) int {
	openNeeded := 0  // 추가로 필요한 여는 괄호의 수
	closeNeeded := 0 // 추가로 필요한 닫는 괄호의 수

	for _, ch := range inputStr {
		if ch == '(' {
			closeNeeded++
		} else { // ch == ')'
			if closeNeeded > 0 {
				closeNeeded--
			} else {
				openNeeded++
			}
		}
	}

	// 최종적으로 필요한 여는 괄호와 닫는 괄호의 수를 더해 반환
	return openNeeded + closeNeeded
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputStr, _ := reader.ReadString('\n')
	inputStr = inputStr[:len(inputStr)-1] // 마지막 개행 문자 제거

	minBracketsToAdd := processBracketSequence(inputStr)
	fmt.Println(minBracketsToAdd)
}
