package main

func main() {

}

func isCousins(root *TreeNode, x int, y int) bool {
	rootx, dx := middleSearch(root, x, 0)
	rooty, dy := middleSearch(root, y, 0)
	if rootx == nil || rooty == nil {
		return false
	}
	if rootx == rooty {
		return false
	}

	return dx == dy
}

func middleSearch(root *TreeNode, v int, k int) (*TreeNode, int) {
	if root == nil {
		return nil, k
	}
	if root.Left != nil && root.Left.Val == v {
		return root, k + 1
	}

	if root.Right != nil && root.Right.Val == v {
		return root, k + 1
	}

	ans,depth := middleSearch(root.Left, v, k+1)
	if ans != nil {
		return ans, depth
	}
	return middleSearch(root.Right, v, k+1)
}

type TreeNode struct {
	  Val int
	  Left *TreeNode
	  Right *TreeNode
}
