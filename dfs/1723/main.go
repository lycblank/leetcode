package main

import (
	"fmt"
	"math"
)

func main() {
	jobs := []int{1,2,4,7,8}
	fmt.Println(minimumTimeRequired(jobs, 2))
}

func minimumTimeRequired(jobs []int, k int) int {
	var ans int = math.MaxInt32
	sum := make([]int, k)
	dfs(jobs, 0, 0, sum, 0, &ans)
	return ans
}

// jobs 工作数组
// u 分配第几个工作
// used 记录已经分配了多少个工人的工作 最要用来求第一个ans使其尽可能小
// max 当前策略下分配到的工作量最大值
// ans 当前最优解
func dfs(jobs []int, u int, used int, sum []int, max int, ans *int) {
	if max >= *ans {
		return
	}
	if u == len(jobs) { // 工作已经被分配完了
		*ans = max
		return
	}

	if used < len(sum) {
		sum[used] = jobs[u]
		dfs(jobs, u + 1, used + 1, sum, MaxInt(sum[used], max), ans)
		sum[used] -= 0
	}

	for i,cnt := 0,used;i<cnt;i++{
		sum[i] += jobs[u]
		dfs(jobs, u + 1, used, sum, MaxInt(sum[i], max), ans)
		sum[i] -= jobs[u]
	}
}

func MaxInt(left int, right int) int {
	if left < right {
		return right
	}
	return left
}

