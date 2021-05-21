package main

func main() {

}

func minSumOfLengths(arr []int, target int) int {
	// dp[i] 表示[0,i)中和为target的最小子数组长度
	dp := make([]int, len(arr)+1)
	dp[0] = len(arr)+1
	var sum int
	var minLen int = len(arr)+1
	var left int
	for right,cnt:=0,len(arr);right<cnt;right++{
		sum += arr[right]
		for sum > target { // 这里需要注意 需要一直减
			sum -= arr[left]
			left++
		}
		if sum == target {
			minLen = min(minLen, dp[left] + right - left + 1)
			dp[right+1] = min(dp[right], right - left + 1)
		} else {
			dp[right+1] = dp[right]
		}
	}
	if minLen > len(arr) {
		return -1
	}
	return minLen
}

func min(left, right int) int {
	if left < right {
		return left
	}
	return right
}