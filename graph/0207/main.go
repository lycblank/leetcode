package main

func main() {

}

func canFinish(numCourses int, prerequisites [][]int) bool {
	degree := make([]int, numCourses)
	edges := make([][]int, numCourses)
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
		degree[pre[0]]++
	}
	queue := make([]int, 0, 8)
	for i := 0; i < numCourses;i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}
	ans := make([]int, 0, numCourses)
	for len(queue) > 0 {
		v := queue[len(queue)-1]
		ans = append(ans, v)
		queue = queue[:len(queue)-1]
		// 删除边
		for _, e := range edges[v] {
			degree[e]--
			if degree[e] <= 0 {
				queue = append(queue, e)
			}
		}
	}
	return len(ans) >= numCourses
}