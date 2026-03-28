package main

import (
	"fmt"
)

func main() {
	fmt.Println(mergeKLists(
		[]*ListNode{
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			}, &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			}, &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 6,
				},
			},
		},
	))
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
func mergeKLists(lists []*ListNode) *ListNode {

	//幾串listNode
	n := len(lists)

}
