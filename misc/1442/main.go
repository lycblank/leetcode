package main

func main() {
	countTriplets([]int{2,3,1,6,7})
}

func countTriplets(arr []int) int {
	sum := make([]int, len(arr)+1)
	for i,cnt:=1,len(arr); i<= cnt;i++{
		sum[i] = sum[i-1]^arr[i-1]
	}

	var ans int
	rets := make(map[int][]int)
	for k,cnt:=0,len(arr);k<=cnt;k++{
		for _, v := range rets[sum[k]] {
			i := v + 1
			ans += k - i
		}
		rets[sum[k]] = append(rets[sum[k]], k)
	}
	return ans
}
