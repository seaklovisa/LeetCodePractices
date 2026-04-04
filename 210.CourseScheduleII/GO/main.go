package main

import "fmt"

func main() {

	prerequisites := [][]int{
		//{1, 0}, {2, 0}, {3, 1}, {3, 2},
		{1, 0}, {1, 2}, {0, 1},
	}

	fmt.Println(findOrder(3, prerequisites))

}

// [1,0],[2,0],[3,1],[3,2]
func findOrder2(numCourses int, prerequisites [][]int) []int {
	//轉換成graph 展示為 課0 <- 課1
	//計算 每個課有的課數

	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)

	for _, p := range prerequisites {
		pre := p[1]
		child := p[0]

		graph[pre] = append(graph[pre], child)
		indegree[child]++
	}

	//把沒有pre的課放到queue內
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	//從queue內取出有pre的課,loop
	order := []int{}

	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]

		order = append(order, course)

		for _, child := range graph[course] {
			indegree[child]--
			if indegree[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	if len(order) != numCourses {
		return []int{}
	}

	return order
}

// [[1,0],[1,2],[0,1]]
// [1,0],[2,0],[3,1],[3,2]
func findOrder(numCourses int, prerequisites [][]int) []int {
	//建圖 (adjancency list)
	graph := make([][]int, numCourses)

	//入度
	indegree := make([]int, numCourses)

	//建立graph+indegree
	for _, p := range prerequisites {
		course := p[0]
		pre := p[1]

		//pre -> course
		graph[pre] = append(graph[pre], course)

		indegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	order := []int{}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		order = append(order, cur)

		for _, next := range graph[cur] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	//無法完成所有課程
	if len(order) != numCourses {
		return []int{}
	}

	return order
}
