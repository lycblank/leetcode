package main

func main() {

}

func numSquares(n int) int {
	nums := make([]int, 100)
	for i:=1;i<=100;i++{
		nums[i-1] = i * i
	}
	// dp[j]表示和j的最小数量
	dp := make([]int, n+1)
	for i:=0;i<=n;i++{
		dp[i] = 1e5
	}
	dp[0] = 0
	for i:=1;i<=100;i++{
		v := nums[i-1]
		if v > n {
			break
		}
		for j := v; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-v]+1)
		}
	}
	return dp[n]
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}
