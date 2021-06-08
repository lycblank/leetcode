package main

func main() {

}

func lastStoneWeightII(stones []int) int {
	sum := 0
	for i := 0; i < len(stones);i++{
		sum += stones[i]
	}
	target := sum / 2
	dp := make([][]int, len(stones)+1)
	for i := 0; i < len(dp);i++{
		dp[i] = make([]int, target+1)
	}

	for i:=1;i<=len(stones);i++{
		for j:=0;j<=target;j++{
			dp[i][j] = dp[i-1][j]
			if j - stones[i-1] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-stones[i-1]]+stones[i-1])
			}
		}
	}

	return sum - dp[len(stones)][target] * 2
}

func max(left, right int) int {
	if left < right {
		return right
	}
	return left
}



