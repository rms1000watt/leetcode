// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// type Node struct {
// 	Val int
// 	Weight int
// 	Neighbors []Node
// }

// func main() {
// 	roads := [][]int{}
// 	roads = append(roads,
// 		[]int{0, 6, 7},
// 		[]int{0, 1, 2},
// 		[]int{1, 2, 3},
// 		[]int{1, 3, 3},
// 		[]int{6, 3, 3},
// 		[]int{3, 5, 1},
// 		[]int{6, 5, 1},
// 		[]int{2, 5, 1},
// 		[]int{0, 4, 5},
// 		[]int{4, 6, 2},
// 	)

// 	countPaths(7, roads)
// }

// // Find all combinations with < n nodes && == shortest
// func countPaths(n int, roads [][]int) (out int) {
// 	graph := createGraph(roads)
// 	fmt.Println(graph)

// 	stack := []Node{}
// 	dfs(n-1, root, graph, stack)

// 	return
// }

// func dfs(destination int, node Node,  graph map[int][]Node, stack []Node) {
// 	for _, neighbor := range node.Neighbors

// }

// // 	destination := n - 1
// // 	for {

// // 		nodes := graph[cursor]
// // 		for _, node := range nodes {
// // 			if visited[]
// // 		}
// // 		currentDistance += node.Weight
// // 		nextNode := graph[node.Neighbor][0]
// // 		currentDistance += nextNode.Weight
// // 		thirdNode := graph

// // 		for _, node := range graph[cursor] {
// // 			currentPath += strconv.Itoa(cursor)

// // 			if !visited[node.Neighbor] {
// // 				cursor = node.Neighbor
// // 				visited[node.Neighbor] = true
// // 				currentDistance += node.Weight
// // 			}
// // 		}
// // 	}

// // 	return
// // }

// func createGraph(roads [][]int) (graph map[int][]Node) {
// 	graph = map[int][]Node{}

// 	for _, road := range roads {
// 		intersection1 := road[0]
// 		intersection2 := road[1]
// 		time := road[2]

// 		graph[intersection1] = append(graph[intersection1], Node{
// 			Val: inters
// 			Weight: time,
// 			Neighbor:   intersection2,
// 		})

// 		graph[intersection2] = append(graph[intersection2], Node{
// 			Weight: time,
// 			Neighbor:   intersection1,
// 		})
// 	}

// 	return
// }
