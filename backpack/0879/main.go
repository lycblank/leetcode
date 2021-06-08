package main

func main() {

}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	dp := make([][]int, n+1)

	sum := 0
	for i := 0; i < len(profit); i++ {
		sum += profit[i]
	}

	for i:=0;i<len(dp);i++{
		dp[i] = make([]int, sum+1)
	}
	dp[0][0] = 1
	for i := 0; i < len(group); i++ {
		for j := n; j >= group[i]; j-- {
			for k := sum; k >= profit[i]; k-- {
				dp[j][k] = (dp[j][k] + dp[j-group[i]][k-profit[i]]) % (1e9+7)
			}
		}
	}
	var ans int
	for k := minProfit; k <= sum; k++ {
		for j:=0; j<=n;j++{
			ans = (ans + dp[j][k]) % (1e9+7)
		}
	}
	return ans
}

