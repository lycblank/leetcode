package main

import "sort"

func main() {

}

func maxSumRangeQuery(nums []int, requests [][]int) int {
	n := len(nums)
	b := make([]int, n)
	for _, q := range requests {
		b[q[0]] += 1
		if q[1] + 1 < n {
			b[q[1]+1] -= 1
		}
	}
	counts := make([]int, n)
	for i, c := range b {
		if i == 0 {
			counts[i] = c
		} else {
			counts[i] = counts[i-1]+c
		}
	}
	sort.Ints(counts)
	sort.Ints(nums)
	var sum int
	for i := n - 1; i >= 0; i-- {
		if counts[i] <= 0 {
			break
		}
		sum += nums[i]*counts[i]
		sum = sum % (1e9+7)
	}
	return sum
}