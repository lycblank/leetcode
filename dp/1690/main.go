package main

func main() {

}

func stoneGameVII(stones []int) int {
	n := len(stones)
	sum := make([]int, n+1)
	for i:=1;i<=n;i++{
		sum[i] = sum[i-1]+stones[i-1]
	}

	dp := make([][]int, n+1)
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int, n+1)
	}

	for i:=n-1;i>=1;i--{
		for j:=i+1;j<=n;j++{
			dp[i][j] = max(sum[j]-sum[i]-dp[i+1][j], sum[j-1]-sum[i-1]-dp[i][j-1])
		}
	}
	return abs(dp[1][n])
}

func max(left, right int) int {
	if left < right {
		return right
	}
	return left
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}


