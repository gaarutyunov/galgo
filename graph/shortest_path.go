package graph

import (
	"github.com/gaarutyunov/galgo/heap"
	"github.com/gaarutyunov/galgo/numbers"
)

func ShortestPathDijkstra(graph [][]int, a, b int) (path []int, length int) {
	edgeWeight := make([]int, len(graph))
	pathWeight := make([]int, len(graph))
	previous := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		pathWeight[i] = numbers.MaxInt
	}
	pathWeight[a] = 0
	remaining := heap.NewMin(pathWeight) // TODO: add key function

	for !remaining.Empty() {
		u, _ := remaining.Pop()
		neighbors := getNeighbors(graph, u)
		for _, v := range neighbors {
			newWeight := pathWeight[u] + edgeWeight[v]
			if newWeight < pathWeight[v] {
				pathWeight[v] = newWeight
				remaining.Set(v, newWeight)
			}
			previous[v] = u
		}
	}

	curr := b

	for previous[curr] != a {
		path = append(path, curr)
		curr = previous[curr]
	}

	length = previous[b]

	return
}

func getNeighbors(graph [][]int, n int) []int {
	panic("not implemented")
}
