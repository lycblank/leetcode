package main

import (
	"math"
	"sort"
)

func main() {

}

func atMostNGivenDigitSet(digits []string, n int) int {
	nums := make([]int, len(digits))
	for i, d := range digits {
		nums[i] = int(d[0]-'0')
	}
	sort.Ints(nums)

	bits := make([]int, 0, 9)
	for n>0{
		bits = append(bits, n % 10)
		n /= 10
	}

	dp := make([]int, len(bits)+1)
	dp[0] = 1
	for i:=1;i<len(dp);i++{
		for _, num := range nums {
			if num < bits[i-1] {
				dp[i] += int(math.Pow(float64(len(nums)), float64(i-1)) + 0.5)
			}
			if num == bits[i-1] {
				dp[i] += dp[i-1]
			}
		}
	}
	ans := dp[len(dp)-1]
	for i:=len(bits)-1;i>=1;i--{
		ans += int(math.Pow(float64(len(nums)), float64(i)) + 0.5)
	}

	return ans
}

