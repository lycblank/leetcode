package main

func main() {

}


// 完全背包问题
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i:=1;i<=len(coins);i++{
		v := coins[i-1]
		for j := v; j <= amount; j++ {
			dp[j] += dp[j-v]
		}
	}
	return dp[amount]
}
