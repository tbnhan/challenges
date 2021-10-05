package main

import "fmt"

// Constraints: 1 <= n <= 3 * 104
const MAX = 312

func main() {
	edges := make([][]int, 0)

	edge1 := []int{0, 1}
	edge2 := []int{0, 2}
	edge3 := []int{2, 3}
	edge4 := []int{2, 4}
	edge5 := []int{2, 5}

	edges = append(edges, edge1)
	edges = append(edges, edge2)
	edges = append(edges, edge3)
	edges = append(edges, edge4)
	edges = append(edges, edge5)

	result := sumOfDistancesInTree(6, edges)
	fmt.Println(result)
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, MAX, MAX)
	for i := 0; i < n-1; i++ {
		u := edges[i][0]
		v := edges[i][1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	result := make([]int, 0, n)

	for i := 0; i < n; i++ {
		path := BFS(i, n, graph)
		sum := 0
		for j := 0; j < n; j++ {
			if j != i {
				distance := getDistance(path, i, j)
				sum = sum + distance
			}
		}
		result = append(result, sum)
	}

	return result
}

func getDistance(path []int, s int, f int) int {
	route := make([]int, 0)

	if s == f {
		return 0
	}

	if path[f] == -1 {
		return 0
	}

	for true {
		route = append(route, f)
		f = path[f]
		if f == s {
			route = append(route, s)
			break
		}
	}

	m := len(route)
	return m - 1
}

func BFS(s int, numOfVertex int, graph [][]int) []int {
	visited := make([]bool, numOfVertex)
	path := make([]int, numOfVertex)
	queue := make([]int, 0)

	visited[s] = true
	queue = append(queue, s)

	for len(queue) != 0 {
		u := queue[0]
		queue = queue[1:]

		numOfNeighbor := len(graph[u])
		for i := 0; i < numOfNeighbor; i++ {
			v := graph[u][i]
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
				path[v] = u
			}
		}
	}

	return path
}
