package main

import "fmt"

func main() {
	nums := []int{1,1,1,1,1}
	fmt.Println(findTargetSumWays(nums, 3))
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i,cnt:=0,len(nums);i<cnt;i++{
		sum += nums[i]
	}
	if target > sum {
		return 0
	}
	dp := make([][]int, len(nums)+1)
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int, sum*2+1)
	}
	dp[0][sum] = 1
	for i := 1; i < len(dp); i++ {
		for j :=-sum; j <= sum; j++ {
			v := nums[i-1]
			if j-v+sum >= 0  { // +
				dp[i][j+sum] += dp[i-1][j-v+sum]
			}
			if j+v+sum <= 2*sum { // -
				dp[i][j+sum] += dp[i-1][j+v+sum]
			}
		}
	}
	return dp[len(nums)][sum+target]
}
