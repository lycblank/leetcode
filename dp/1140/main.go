package main

func main() {

}

func stoneGameII(piles []int) int {
	n := len(piles)
	// dp表示[i:len]中，M为j时，先手能取到的最大石头数量
	dp := make([][]int, n+1)
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int, n+1)
	}
	sum := 0
	for i := n; i>=1;i--{
		sum += piles[i-1]
		for j:=1;j<=n;j++{
			if i-1+2*j>=n{
				dp[i][j] = sum
			} else {
				for m:=1;m<=2*j;m++{
					dp[i][j] = max(dp[i][j], sum - dp[i+m][max(m,j)])
				}
			}
		}
	}
	return dp[1][1]
}

func max(left, right int) int {
	if left < right {
		return right
	}
	return left
}