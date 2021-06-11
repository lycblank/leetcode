package main

func main() {

}

func stoneGameVIII(stones []int) int {
	n := len(stones)
	sum := make([]int, n+1)
	for i:=1;i<=n;i++{
		sum[i] = sum[i-1]+stones[i-1]
	}
	// dp[i]表示从i:n石子差
	dp := make([]int, n+1)
	dp[n] = sum[n]
	for i:=n-1;i>=1;i--{
		// 不选i => dp[i+1]
		// 选i=> sum[i]-dp[i+1]
		dp[i] = max(dp[i+1], sum[i] - dp[i+1])
	}
	return dp[2]
}

func max(left,right int) int {
	if left < right {
		return right
	}
	return left
}

