// TODO: 버블 정렬 (Bubble Sort)
// TODO: 서로 인접한 두 원소를 검사하여 정렬

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
	inputString := scanner.Text()

	bubbleSortItems := strings.Split(inputString, " ")
	numbers := make([]int, len(bubbleSortItems))

	for i, str := range bubbleSortItems {
		number, _ := strconv.Atoi(str)
		numbers[i] = number
	}

	bubbleSort(numbers)
	fmt.Println(numbers)
}

func bubbleSort(numbers []int) {
	n := len(numbers)

	for i := 0; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}
