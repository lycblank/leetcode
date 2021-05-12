package main

func main() {

}

func xorQueries(arr []int, queries [][]int) []int {
	ans := make([]int, len(queries))
	sum := make([]int, len(arr))
	val := 0
	for i,cnt:=0,len(arr);i<cnt;i++{
		val = val ^ arr[i]
		sum[i] = val
	}
	for i,cnt:=0,len(queries);i<cnt;i++{
		if queries[i][0] == 0 {
			ans[i] = sum[queries[i][1]]
		} else {
			ans[i] = sum[queries[i][0]-1] ^ sum[queries[i][1]]
		}
	}
	return ans
}