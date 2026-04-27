package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}

	connect(root)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	leftMost := root

	for leftMost.Left != nil {
		cur := leftMost

		for cur != nil {
			cur.Left.Next = cur.Right

			if cur.Next != nil {
				cur.Right.Next = cur.Next.Left
			}

			cur = cur.Next
		}
		leftMost = leftMost.Left
	}

	return root
}
