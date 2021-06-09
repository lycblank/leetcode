package main

func main() {

}

func stoneGame(piles []int) bool {
	// dp[i][j]表示先手比后手多最大的石子
	dp := make([][]int, len(piles)+1)
	for i:=0;i<len(dp);i++ {
		dp[i] = make([]int, len(piles)+1)
	}
	for i:=1;i<len(dp);i++ {
		dp[i][i] = piles[i-1]
	}
	for i:=len(piles)-1;i>=1;i-- {
		for j:=i+1;j<=len(piles);j++{
			dp[i][j] = max(piles[i-1] - dp[i+1][j], piles[j-1]-dp[i][j-1])
		}
	}
	return dp[1][len(piles)] > 0
}

func max(left int, right int) int {
	if left < right {
		return right
	}
	return left
}
