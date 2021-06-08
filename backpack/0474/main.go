package main

func main() {

}

func findMaxForm(strs []string, m int, n int) int {
	one := make([]int, len(strs))
	zero := make([]int, len(strs))
	for i:=0;i<len(strs);i++{
		for j:=0;j<len(strs[i]);j++{
			if strs[i][j] == '0' {
				zero[i]++
			} else {
				one[i]++
			}
		}
	}
	dp := make([][]int, m+1)
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int, n+1)
	}
	for i:=0;i<len(strs);i++{
		for j := m; j >= zero[i]; j-- {
			for k := n; k >= one[i]; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zero[i]][k-one[i]]+1)
			}
		}
	}
	return dp[m][n]
}

func max(left, right int) int {
	if left < right {
		return right
	}
	return left
}
