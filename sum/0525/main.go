package main

func main() {

}

func findMaxLength(nums []int) int {
	var tmp, ans int
	others := make(map[int]int)
	others[0] = -1
	for i,cnt:=0,len(nums);i<cnt;i++{
		if nums[i] == 0 {
			tmp += -1
		} else {
			tmp += 1
		}
		idx, ok := others[tmp]
		if ok && i - idx > ans {
			ans = i - idx
		}
		if !ok {
			others[tmp] = i
		}
	}
	return ans
}

