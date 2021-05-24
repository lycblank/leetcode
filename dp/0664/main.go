package main

import "math"

func main() {

}

func strangePrinter(s string) int {
	dp := make([][]int, len(s))
	for i,cnt:=0,len(dp);i<cnt;i++{
		dp[i] = make([]int, len(s))
		for j,cnt1:=0,len(dp[i]);j<cnt1;j++{
			dp[i][j] = math.MaxInt64
		}
		dp[i][i] = 1
	}
	for i:=len(dp)-1;i>=0;i--{
		for j:=i+1;j<len(dp[i]);j++{
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
				if i - 1 > 0 && dp[i][j] > dp[i-1][j] {
					dp[i][j] = dp[i-1][j]
				}
				continue
			}
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k] + dp[k+1][j])
			}
		}
	}
	return dp[0][len(s)-1]
}

func min(left,right int) int {
	if left < right {
		return left
	}
	return right
}

