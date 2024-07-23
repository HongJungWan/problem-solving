// TODO: 소수 찾기, 순열
// https://school.programmers.co.kr/learn/courses/30/lessons/42839

package main

import (
	"fmt"
	"math"
	"strconv"
)

// 소수 판별 함수
func isPrimeTwo(n int) bool {
	if n <= 1 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 문자열의 모든 순열을 찾는 함수
func generatePermutations(numbers string, current string, visited []bool, result *map[int]bool) {
	if current != "" {
		num, _ := strconv.Atoi(current)
		(*result)[num] = true
	}
	for i := 0; i < len(numbers); i++ {
		if !visited[i] {
			visited[i] = true
			generatePermutations(numbers, current+string(numbers[i]), visited, result)
			visited[i] = false
		}
	}
}

// 주어진 숫자 문자열에서 조합 가능한 소수의 개수를 찾는 함수
func countPrimes(numbers string) int {
	result := make(map[int]bool)
	visited := make([]bool, len(numbers))
	generatePermutations(numbers, "", visited, &result)

	primeCount := 0
	for num := range result {
		if isPrimeTwo(num) {
			primeCount++
		}
	}
	return primeCount
}

func main() {
	numbers := "17"
	fmt.Printf("소수의 개수: %d\n", countPrimes(numbers)) // 3

	numbers = "011"
	fmt.Printf("소수의 개수: %d\n", countPrimes(numbers)) // 2
}

/*
[순열] :
- 순열(Permutation)은 주어진 집합에서 순서를 고려하여 요소를 배치하는 방법
- 예를 들어, 집합 {1, 2, 3}의 모든 순열은 {1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}


[DFS를 사용해 순열을 구하는 연산 과정] :
예제: numbers = "17"

- 초기 상태
  - current = ""
  - visited = [false, false]
  - result = map[int]bool{}

- 첫 번째 재귀 호출
  - current = ""
  - 각 자리 숫자에 대해:
    - i = 0 (숫자 '1')
      - visited = [true, false]
      - current = "1"
      - 재귀적으로 generatePermutations(numbers, "1", visited, result) 호출

- 두 번째 재귀 호출 (current = "1")
  - current = "1"
  - current가 비어 있지 않으므로, 1을 숫자로 변환하여 result에 추가: result = map[int]bool{1: true}
  - 각 자리 숫자에 대해:
    - i = 0 (이미 방문했으므로 패스)
    - i = 1 (숫자 '7')
      - visited = [true, true]
      - current = "17"
      - 재귀적으로 generatePermutations(numbers, "17", visited, result) 호출

- 세 번째 재귀 호출 (current = "17")
  - current = "17"
  - current가 비어 있지 않으므로, 17을 숫자로 변환하여 result에 추가: result = map[int]bool{1: true, 17: true}
  - 각 자리 숫자에 대해:
    - i = 0 (이미 방문했으므로 패스)
    - i = 1 (이미 방문했으므로 패스)
  - 재귀 호출 종료, 백트래킹: visited = [true, false], current = "1"

- 백트래킹 후 상태 (current = "1")
  - visited = [true, false]
  - current = "1"
  - i = 1 (숫자 '7')
    - visited = [true, true]
    - current = "17"
    - 재귀적으로 generatePermutations(numbers, "17", visited, result) 호출
  - 이미 result에 "17"이 있으므로 중복된 값을 추가하지 않음
  - 재귀 호출 종료, 백트래킹: visited = [false, false], current = ""

- 첫 번째 재귀 호출로 돌아와서 (current = "")
  - i = 1 (숫자 '7')
    - visited = [false, true]
    - current = "7"
    - 재귀적으로 generatePermutations(numbers, "7", visited, result) 호출
- 두 번째 재귀 호출 (current = "7")
  - current = "7"
  - current가 비어 있지 않으므로, 7을 숫자로 변환하여 result에 추가: result = map[int]bool{1: true, 17: true, 7: true}
  - 각 자리 숫자에 대해:
    - i = 0 (숫자 '1')
      - visited = [true, true]
      - current = "71"
      - 재귀적으로 generatePermutations(numbers, "71", visited, result) 호출

- 세 번째 재귀 호출 (current = "71")
  - current = "71"
  - current가 비어 있지 않으므로, 71을 숫자로 변환하여 result에 추가: result = map[int]bool{1: true, 17: true, 7: true, 71: true}
  - 각 자리 숫자에 대해:
    - i = 0 (이미 방문했으므로 패스)
    - i = 1 (이미 방문했으므로 패스)
  - 재귀 호출 종료, 백트래킹: visited = [false, true], current = "7"

- 모든 순열 생성 후 result 상태
  - result = map[int]bool{1: true, 17: true, 7: true, 71: true}

- 소수 판별 및 개수 세기
  - 1은 소수가 아니므로 제외
  - 17은 소수
  - 7은 소수
  - 71은 소수

- 따라서 소수의 개수는 3
- 최종 결과
  - 소수의 개수는 3
*/
