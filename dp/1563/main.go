package main

func main() {

}

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	sum := make([]int,n+1)
	for i:=1;i<=n;i++{
		sum[i] = sum[i-1]+stoneValue[i-1]
	}
	// dp表示i到j之间能得到的最高分
	dp := make([][]int, n+1)
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int, n+1)
	}
	for i:=n;i>=1;i--{
		for j:=i+1;j<=n;j++{
			for k:=i;k<j;k++{
				left := sum[k]-sum[i-1]
				right := sum[j]-sum[k]
				if left < right {
					dp[i][j] = max(dp[i][j], left + dp[i][k])
				} else if left > right {
					dp[i][j] = max(dp[i][j], right + dp[k+1][j])
				} else {
					dp[i][j] = max(dp[i][j], left + max(dp[i][k],dp[k+1][j]))
				}
			}
		}
	}
	return dp[1][n]
}

func max(left,right int) int {
	if left < right {
		return right
	}
	return left
}


