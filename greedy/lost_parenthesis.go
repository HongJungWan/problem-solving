// TODO: 잃어버린 괄호
// https://www.acmicpc.net/problem/1541

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var expression string
	fmt.Scanln(&expression)

	// '-'로 분할
	parts := strings.Split(expression, "-")
	result := 0

	for i, part := range parts {
		// '+'로 분할하여 합을 구함
		sum := 0
		for _, num := range strings.Split(part, "+") {
			val, _ := strconv.Atoi(num)
			sum += val
		}

		// 첫 부분은 더하고, 그 다음부터는 뺌
		if i == 0 {
			result += sum
		} else {
			result -= sum
		}
	}

	fmt.Println(result)
}

// -로 분할하고, 첫 번째 항목을 제외한 나머지를 모두 빼는 것이 괄호를 통해 최소값을 얻는 공식
