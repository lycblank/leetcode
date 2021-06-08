package main

func main() {

}

func waysToChange(n int) int {
	coins := []int{1,5,10, 25}
	dp := make([]int, n+1)
	dp[0] = 1
	for _, c := range coins {
		for i := c; i <= n; i++{
			dp[i] = (dp[i] + dp[i-c])%1000000007
		}
	}
	return dp[n]
}
