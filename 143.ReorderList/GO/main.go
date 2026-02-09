package main

func main() {
	head3 := &ListNode{
		Val:  1,
		Next: nil,
	}

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	// head2 := &ListNode{
	// 	Val: 1,
	// 	Next: &ListNode{
	// 		Val: 2,
	// 		Next: &ListNode{
	// 			Val: 3,
	// 			Next: &ListNode{
	// 				Val: 4,
	// 				Next: &ListNode{
	// 					Val:  5,
	// 					Next: nil,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	reorderList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func recordList2(head *ListNode) {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {

	//從中間斷開
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil

	//反轉第二段
	var prev *ListNode
	cur := second

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	second = prev

	//交叉重組
	first := head

	for second != nil {
		t1 := first.Next
		t2 := second.Next

		first.Next = second
		second.Next = t1

		first = t1
		second = t2
	}
}
