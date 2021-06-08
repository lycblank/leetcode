package main

import (
	"container/heap"
)

func main() {

}

func lastStoneWeight(stones []int) int {
	h := Heap(stones)
	heap.Init(&h)
	for len(h) > 1 {
		v1 := heap.Pop(&h).(int)
		v2 := heap.Pop(&h).(int)
		diff := v1 - v2
		if diff > 0 {
			heap.Push(&h, diff)
		}
	}
	if len(h) > 0 {
		return h[0]
	}
	return 0
}

type Heap []int
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{}  {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
