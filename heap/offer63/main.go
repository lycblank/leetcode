package main

import "container/heap"

func main() {

}

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	buyList := Heap(make([]int, 0, len(prices)))
	heap.Push(&buyList, prices[0])
	var ans int
	for i:=1;i<len(prices);i++{
		ans = max(ans, prices[i]-buyList[0])
		heap.Push(&buyList, prices[i])
	}
	return ans
}

func max(left, right int) int {
	if left < right {
		return right
	}
	return left
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
	return (*h)[i] < (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

