package main

// https://leetcode.com/problems/clone-graph/submissions/

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) (out *Node) {
	graph := make(map[int]*Node)

	// clone the first node
	// clone will recursively be called and clone all neighbors
	// but you want to return the first node
	out = clone(node, graph)

	return
}

func clone(node *Node, graph map[int]*Node) (out *Node) {
	if node == nil {
		return
	}

	if graph[node.Val] == nil {
		graph[node.Val] = &Node{
			Val: node.Val,
		}

		for _, neighbor := range node.Neighbors {
			graph[node.Val].Neighbors = append(graph[node.Val].Neighbors, clone(neighbor, graph))
		}
	}

	out = graph[node.Val]
	return
}
