package main

import "math/rand"

func main() {

}

type Solution struct {
	head *ListNode
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{
		head:head,
	}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	var ans int
	var i int
	for next := this.head; next != nil; next = next.Next {
		if rand.Intn(i+1) < 1 {
			ans = next.Val
		}
		i++
	}
	return ans
}

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
