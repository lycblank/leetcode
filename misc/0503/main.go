package main

func main() {

}

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	stack := []int{}
	for i,cnt:=0,2*n; i < cnt; i++ {
		idx := i % n
		var j int
		for j = len(stack) - 1; j >= 0; j-- {
			if nums[stack[j]] < nums[idx] {
				ans[stack[j]] = nums[idx]
				continue
			}
			break
		}
		stack = stack[:j+1]
		stack = append(stack, idx)
	}
	return ans
}