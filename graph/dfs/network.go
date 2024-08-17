// 네트워크
// https://school.programmers.co.kr/learn/courses/30/lessons/43162

package main

import "fmt"

func networkDFS(computers [][]int, visited []bool, current int) {
	visited[current] = true
	for i, connection := range computers[current] {
		if connection == 1 && !visited[i] {
			networkDFS(computers, visited, i)
		}
	}
}

func countNetworks(n int, computers [][]int) int {
	visited := make([]bool, n)
	count := 0
	for i := range visited {
		if !visited[i] {
			networkDFS(computers, visited, i)
			count++
		}
	}
	return count
}

func main() {
	n := 3
	computers := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	fmt.Println(countNetworks(n, computers)) // 2
}
