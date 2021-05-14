package main

func main() {

}

func countRangeSum(nums []int, lower int, upper int) int {
	sum := make([]int, len(nums)+1)
	sum[0] = nums[0]
	for i,cnt:=0,len(nums);i<cnt;i++{
		sum[i+1] = sum[i] + nums[i]
	}

	return  mergeCountRangeSum(sum, lower, upper)
}

func mergeCountRangeSum(arr []int, lower int, upper int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	n1 := append([]int{}, arr[:n/2]...)
	n2 := append([]int{}, arr[n/2:]...)
	ans := mergeCountRangeSum(n1, lower, upper) + mergeCountRangeSum(n2, lower, upper)
	ans += countSortRangeSum(n1, n2, lower, upper)

	// 合并排序
	var i1, i2, idx int
	for ;idx<n&&i1<len(n1)&&i2<len(n2);idx++{
		if n1[i1] < n2[i2] {
			arr[idx] = n1[i1]
			i1++
			continue
		}
		arr[idx] = n2[i2]
		i2++
	}

	if i1 >= len(n1) {
		for ;i2<len(n2);i2++{
			arr[idx] = n2[i2]
			idx++
		}
	}

	if i2 >= len(n2) {
		for ;i1<len(n1);i1++{
			arr[idx] = n1[i1]
			idx++
		}
	}

	return ans
}

func countSortRangeSum(arr1 []int, arr2 []int, lower int, upper int) int {
	var l, r, ans int
	for _, val := range arr1 {
		for ;l<len(arr2);l++{
			if arr2[l] >= lower+val {
				break
			}
		}
		for ;r<len(arr2);r++{
			if arr2[r]>upper+val {
				break
			}
		}
		ans += r - l
	}
	return ans
}
