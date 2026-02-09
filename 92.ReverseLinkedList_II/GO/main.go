package main

import "fmt"

func main() {

	fmt.Println(reverseBetween(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}, 2, 4))

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
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head.Next == nil || left == right {
		return head
	}

	dummy := &ListNode{Next: head}

	pre := dummy

	//走到left前
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	//開始反轉
	cur := pre.Next

	for i := 0; i < right-left; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp

	}

	return dummy.Next

}
