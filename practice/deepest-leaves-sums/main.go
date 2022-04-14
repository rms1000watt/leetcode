// https://leetcode.com/problems/deepest-leaves-sum/submissions/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	deepestHashTable map[int][]int
	deepestLevel     int
)

func deepestLeavesSum(root *TreeNode) (out int) {
	deepestHashTable = map[int][]int{}
	deepestLevel = 0

	dig(0, root)

	for _, v := range deepestHashTable[deepestLevel] {
		out += v
	}

	return
}

func dig(depth int, node *TreeNode) {
	if node == nil {
		return
	}

	if depth > deepestLevel {
		deepestLevel = depth
	}

	if node.Left == nil && node.Right == nil {
		deepestHashTable[depth] = append(deepestHashTable[depth], node.Val)
		return
	}

	depth++
	dig(depth, node.Right)
	dig(depth, node.Left)

	return
}
