package main

import "container/list"

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1, s2 := &Stack{}, &Stack{}
	if root1 != nil {
		s1.Push(root1)
	}
	if root2 != nil {
		s2.Push(root2)
	}
	for !s1.IsEmpty() && !s2.IsEmpty() {
		node1 := s1.Pop().(*TreeNode)
		node2 := s2.Pop().(*TreeNode)
		for node1.Left != nil || node1.Right != nil {
			if node1.Left != nil {
				if node1.Right != nil {
					s1.Push(node1.Right)
				}
				node1 = node1.Left
			} else {
				node1 = node1.Right
			}
		}

		for node2.Left != nil || node2.Right != nil {
			if node2.Left != nil {
				if node2.Right != nil {
					s2.Push(node2.Right)
				}
				node2 = node2.Left
			} else {
				node2 = node2.Right
			}
		}
		if node1.Val != node2.Val {
			return false
		}
	}
	return s1.IsEmpty() && s2.IsEmpty()
}

type Stack struct {
	l list.List
}

func (s *Stack) IsEmpty() bool {
	return s.l.Len() <= 0
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	e := s.l.Back()
	return s.l.Remove(e)
}

func (s *Stack) Push(v interface{}) {
	s.l.PushBack(v)
}




type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}