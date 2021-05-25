package main

import "math"

func main() {

}

func minChanges(nums []int, k int) int {
	maxLen := 1 << 10
	dp := make([][]int, k)
	for i,cnt:=0,len(dp);i<cnt;i++{
		dp[i] = make([]int, maxLen)
	}
	n := len(nums)

	preMin := make([]int, k)

	for i := 0; i < k; i++ {
		preMin[i] = math.MaxInt64
		// 统计第i列的数字
		var cnt int
		cnts := map[int]int{}
		for j := i; j < n; j += k {
			cnts[nums[j]]++
			cnt++
		}
		if i == 0 {
			for v := 0; v < maxLen; v++ {
				dp[i][v] = cnt - cnts[v]
				preMin[i] = min(preMin[i], dp[i][v])
			}
		} else {
			for v := 0; v < maxLen; v++ {
				dp[i][v] = preMin[i-1] + cnt // 整列替换
				for num, c := range cnts { // 部分替换
					dp[i][v] = min(dp[i][v], dp[i-1][num^v]+cnt-c)
				}
				preMin[i] = min(preMin[i], dp[i][v])
			}
		}
	}
	return dp[k-1][0]
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

