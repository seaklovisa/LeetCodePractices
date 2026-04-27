package main

import (
	"math"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

type NodeInfo struct {
	row int
	col int
	val int
}

func verticalTraversal(root *TreeNode) [][]int {
	nodes := []NodeInfo{}

	var dfs func(node *TreeNode, row, col int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}

		nodes = append(nodes, NodeInfo{row, col, node.Val})

		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}

	dfs(root, 0, 0)

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].col != nodes[j].col {
			return nodes[i].col < nodes[j].col
		}
		if nodes[i].row != nodes[j].row {
			return nodes[j].row < nodes[j].row
		}

		return nodes[i].val < nodes[j].val
	})

	res := [][]int{}
	currCol := math.MinInt32
	var colList []int

	for _, n := range nodes {
		if n.col != currCol {
			if len(colList) > 0 {
				res = append(res, colList)
			}
			colList = []int{}
			currCol = n.col
		}
		colList = append(colList, n.val)
	}

	if len(colList) > 0 {
		res = append(res, colList)
	}

	return res
}
