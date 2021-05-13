package main

func main() {

}

func numWays(steps int, arrLen int) int {
	// dp[i][j] 表示有i步停留在j的位置方案数
	const mod = 1e9 + 7
	dp := make([][]int, steps+1)

	edge := arrLen
	if edge > steps {
		edge = steps + 1
	}

	for i:=0;i<steps+1;i++{
		dp[i] = make([]int, edge)
	}
	dp[0][0] = 1

	for i:=1;i<steps+1;i++{
		for j:=0;j<edge;j++{
			dp[i][j] = dp[i-1][j]
			if j - 1 >= 0 {
				dp[i][j] = (dp[i][j]+dp[i-1][j-1])%mod
			}
			if j + 1 < edge {
				dp[i][j] = (dp[i][j]+dp[i-1][j+1])%mod
			}
		}
	}
	return dp[steps][0]
}




