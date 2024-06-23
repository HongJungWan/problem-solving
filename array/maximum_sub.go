// TODO: 최대 부분 합
// TODO: Kadane's Algorithm, 최대 부분 배열의 합을 찾는 'maxSubArray' 함수 사용

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	numString := strings.Split(scanner.Text(), " ")

	nums := make([]int, T)
	for i := 0; i < T; i++ {
		nums[i], _ = strconv.Atoi(numString[i])
	}

	fmt.Printf("%d ", maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}

/*
[입력]
5 (배열의 크기)
2 1 -2 3 -5 (element)

[출력]
4
*/
