package main

func main() {

}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1)+1)
	for i,cnt:=0,len(dp);i<cnt;i++{
		dp[i] = make([]int, len(nums2)+1)
	}
	for i,cnt:=0,len(nums1)+1;i<cnt;i++{
		dp[i][0] = 0
	}
	for j,cnt:=0,len(nums2)+1;j<cnt;j++{
		dp[0][j] = 0
	}
	for i,cnt1 := 0,len(nums1);i<cnt1;i++{
		for j,cnt2:=0,len(nums2);j<cnt2;j++{
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j]+1
				continue
			}
			dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
		}
	}
	return dp[len(nums1)][len(nums2)]
}

func max(left,right int) int {
	if left < right {
		return right
	}
	return left
}