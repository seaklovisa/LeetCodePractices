package main

import (
	"container/list"
	"fmt"
)

func main() {

	fmt.Println(forLoopRecure(5))
}

func forLoopRecure(n int) int {
	stack := list.New()
	res := 0

	for i := n; i > 0; i-- {
		stack.PushBack(i)
	}

	for stack.Len() != 0 {
		res += stack.Back().Value.(int)
		stack.Remove(stack.Back())
	}

	return res
}
