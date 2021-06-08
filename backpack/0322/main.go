package main

import "math"

func main() {

}

func coinChange(coins []int, amount int) int {
	maxInt := math.MaxInt64 -1
	dp := make([]int, amount + 1)
	for i:=0;i<len(dp);i++{
		dp[i] = maxInt
	}
	dp[0] = 0
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] == maxInt {
		return -1
	}
	return dp[amount]
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}

