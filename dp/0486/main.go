package main

func main() {

}

func PredictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums)+1)
	for i:=0;i<len(dp);i++{
		dp[i] = make([]int, len(nums)+1)
	}
	for i:=1;i<=len(nums);i++ {
		dp[i][i] = nums[i-1]
	}
	for i:=len(nums)-1;i>=1;i--{
		for j:=i+1;j<=len(nums);j++{
			dp[i][j] = max(nums[i-1]-dp[i+1][j],nums[j-1]-dp[i][j-1])
		}
	}
	return dp[1][len(nums)] >= 0
}

func max(left,right int) int {
	if left < right {
		return right
	}
	return left
}