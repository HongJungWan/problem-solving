// 바이러스
// https://www.acmicpc.net/problem/2606

package main

import (
	"fmt"
)

var networkGraph map[int][]int
var visitedNetwork []bool
var infectedCount int

func SpreadVirus(startComputer int) {
	visitedNetwork[startComputer] = true
	infectedCount++

	for _, connectedComputer := range networkGraph[startComputer] {
		if !visitedNetwork[connectedComputer] {
			SpreadVirus(connectedComputer)
		}
	}
}

func main() {
	var numComputers, numConnections int
	fmt.Scanf("%d", &numComputers)
	fmt.Scanf("%d", &numConnections)

	networkGraph = make(map[int][]int)
	visitedNetwork = make([]bool, numComputers+1)

	var computer1, computer2 int
	for i := 0; i < numConnections; i++ {
		fmt.Scanf("%d %d", &computer1, &computer2)
		networkGraph[computer1] = append(networkGraph[computer1], computer2)
		networkGraph[computer2] = append(networkGraph[computer2], computer1)
	}

	SpreadVirus(1)
	fmt.Println(infectedCount - 1)
}
