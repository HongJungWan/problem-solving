// TODO: 팰린드롬 체크 문제
// TODO: 어떤 단어를 뒤에서부터 읽어도 똑같다면 그 단어를 팰린드롬

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

	fmt.Println(isPalindrome(inputString))
}

func isPalindrome(inputString string) bool {
	inputStringLength := len(inputString)

	for i := 0; i < inputStringLength/2; i++ {
		if inputString[i] != inputString[inputStringLength-1-i] {
			return false
		}
	}
	return true
}
