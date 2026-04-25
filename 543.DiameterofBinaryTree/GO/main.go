package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	})
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 取得左子樹的值
		left := dfs(node.Left)
		// 取得右子樹的值
		right := dfs(node.Right)

		if left+right > maxDiameter {
			maxDiameter = left + right
		}

		if left > right {
			return left + 1
		}
		return right + 1
	}

	dfs(root)

	return maxDiameter

}
