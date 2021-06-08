package main

func main() {

}

func canPartition(nums []int) bool {
	sum := 0
	for i:=0;i<len(nums);i++{
		sum += nums[i]
	}
	if sum&1 != 0 {
		return false
	}
	sum >>= 1
	dp := make([]bool, sum+1)
	dp[0] = true
	for i:=1;i<=len(nums);i++{
		v := nums[i-1]
		for j:=sum;j>=v;j--{
			dp[j]=dp[j]||dp[j-v]
		}
	}
	return dp[sum]
}

