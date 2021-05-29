package main

func main() {

}

func subarraySum(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	for i:=1;i<len(sum);i++{
		sum[i] = sum[i-1] + nums[i-1]
	}
	var ans int
	others := map[int]int{}
	for i:=0;i<len(sum);i++{
		target := sum[i] - k
		if cnt, ok := others[target]; ok {
			ans += cnt
		}
		others[sum[i]]++
	}
	return ans
}
