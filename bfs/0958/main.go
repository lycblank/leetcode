package main

func main() {

}

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}
func isCompleteTree(root *TreeNode) bool {
	bfs := make([]*TreeNode, 0, 100)
	bfs = append(bfs, root)
	var i int
	for i = 0; i < len(bfs) && bfs[i] != nil; i++ {
		bfs = append(bfs, bfs[i].Left, bfs[i].Right)
	}
	for ;i < len(bfs) && bfs[i] == nil;i++ {
	}
	return i == len(bfs)
}


