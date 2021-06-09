package main
func main() {

}

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	ans := nums[0]
	sum := 0
	if nums[0] > 0 {
		sum = nums[0]
	}
	for i:=1;i<len(nums);i++{
		sum += nums[i]
		if ans < sum {
			ans = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}
	return ans
}
