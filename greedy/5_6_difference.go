// TODO: 5와 6의 차이
// https://www.acmicpc.net/problem/2864

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateSum(A, B string, oldChar, newChar rune) int {
	adjustedA := strings.ReplaceAll(A, string(oldChar), string(newChar))
	adjustedB := strings.ReplaceAll(B, string(oldChar), string(newChar))

	intA, _ := strconv.Atoi(adjustedA)
	intB, _ := strconv.Atoi(adjustedB)

	return intA + intB
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	parts := strings.Split(input, " ")
	A, B := parts[0], parts[1]

	minSum := calculateSum(A, B, '6', '5')
	maxSum := calculateSum(A, B, '5', '6')

	fmt.Println(minSum, maxSum)
}
