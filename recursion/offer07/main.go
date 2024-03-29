package main

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	root := &TreeNode{
		Val:preorder[0],
	}
	var i int
	for i = 0;i<len(inorder);i++{
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

type TreeNode struct {
	  Val int
	  Left *TreeNode
	  Right *TreeNode
}
