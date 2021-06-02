package main

func main() {

}

func checkSubarraySum(nums []int, k int) bool {
	others := make(map[int]int)
	others[0] = -1
	var tmp int
	for i,cnt:=0,len(nums); i < cnt; i++ {
		tmp = (tmp + nums[i]) % k
		idx, ok := others[tmp]
		if ok && i - idx >= 2 {
			return true
		}
		if !ok {
			others[tmp] = i
		}
	}
	return false
}
