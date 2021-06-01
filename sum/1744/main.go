package main

import "fmt"

func main() {
	candiesCount := []int{7,4,5,3,8}
	querys := [][]int{[]int{0,2,2}}
	fmt.Println(canEat(candiesCount, querys))
}

func canEat(candiesCount []int, queries [][]int) []bool {
	sum := make([]int, len(candiesCount)+1)
	for i:=1;i<len(sum);i++ {
		sum[i] = sum[i-1] + candiesCount[i-1]
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		target, day, dailyCap := q[0], q[1]+1, q[2]
		x1 := day
		y1 := dailyCap * day

		x2 := sum[target]+1
		y2 := sum[target+1]

		ans[i] = !(x1 > y2 || x2 > y1)
	}
	return ans
}



