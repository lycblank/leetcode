package main

import "fmt"

func main() {
	arr := []int{0,1,2,3,0}
	fmt.Println(minSideJumps(arr))
}

func minSideJumps(obstacles []int) int {
	dp := make([]int, 3)
	dp[0],dp[1],dp[2] = 1, 0, 1
	for i:=1;i<len(obstacles);i++{
		p1,p2,p3:=dp[0],dp[1],dp[2]
		initDP(dp)
		if obstacles[i] != 1 {
			dp[0] = p1
		}
		if obstacles[i] != 2 {
			dp[1] = p2
		}
		if obstacles[i] != 3 {
			dp[2] = p3
		}
		if obstacles[i] != 1 {
			dp[0] = minInt(dp[0], minInt(dp[1],dp[2])+1)
		}
		if obstacles[i] != 2 {
			dp[1] = minInt(dp[1], minInt(dp[0],dp[2])+1)
		}
		if obstacles[i] != 3 {
			dp[2] = minInt(dp[2], minInt(dp[0],dp[1])+1)
		}
	}
	return minInt(dp[0], minInt(dp[1], dp[2]))
}

func initDP(dp []int) {
	for i:=0;i<len(dp);i++{
		dp[i] = 1e8
	}
}

func minInt(left,right int) int {
	if left < right {
		return left
	}
	return right
}



