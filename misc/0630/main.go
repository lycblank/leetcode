package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		if courses[i][1] < courses[j][1] {
			return true
		}
		if courses[i][1] == courses[j][1] {
			return courses[i][0] < courses[j][0]
		}
		return false
	})
	t := 0
	q := IntHeap{}
	for i := 0; i < len(courses); i++ {
		cost := courses[i][0]
		heap.Push(&q, cost)
		diff := cost
		if t + cost > courses[i][1] {
			maxCost := heap.Pop(&q).(int)
			diff = cost - maxCost
		}
		t += diff
	}
	return len(q)
}

type IntHeap []int
func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}



