package main

import "fmt"

func main() {
	fmt.Println(swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{
		Next: head,
	}

	pre := dummy

	for pre.Next != nil && pre.Next.Next != nil {
		cur := pre.Next
		temp := cur.Next

		//節點交換
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp

		//移動兩次
		pre = pre.Next.Next

	}

	return dummy.Next
}
