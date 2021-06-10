package main

func main() {

}

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	// 表示先手在i:n-1之间拿 能拿到的最大分
	dp := make([]int, n+1)
	sum := 0
	for i:=n-1;i>=0;i--{
		sum += stoneValue[i]
		dp[i] = sum - dp[i+1]
		for j:=i+2;j<=i+3&&j<=n;j++{
			dp[i] = max(dp[i], sum-dp[j])
		}
	}
	if sum - dp[0] < dp[0] { //先手赢
		return "Alice"
	}

	if sum - dp[0] > dp[0] { // 后手赢
		return "Bob"
	}

	return "Tie"
}

func max(left, right int) int {
	if left < right {
		return right
	}
	return left
}

func min(left,right int) int {
	if left < right {
		return left
	}
	return right
}


