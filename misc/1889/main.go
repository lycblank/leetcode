package main

import (
	"math"
	"sort"
)

func main() {

}


func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	sumPkg := make([]int64, len(packages)+1)
	for i:=1;i<=len(packages);i++{
		sumPkg[i] = sumPkg[i-1]+int64(packages[i-1])
	}

	ans := int64(math.MaxInt64)
	for i:=0;i<len(boxes);i++ {
		sort.Ints(boxes[i])
		startIdx := 0
		sum := int64(0)
		if boxes[i][len(boxes[i])-1] < packages[len(packages)-1] {
			continue
		}
		for k := 0; k < len(boxes[i]); k++ {
			if boxes[i][k] < packages[startIdx] {
				continue
			}
			// idx := sort.SearchInts(packages, boxes[i][k])
			idx := sort.Search(len(packages), func(t int) bool {
				return packages[t] > boxes[i][k]
			})

			sum = (sum + int64(idx - startIdx) * int64(boxes[i][k]) - sumPkg[idx] + sumPkg[startIdx]) % (1000000007)
			startIdx = idx
			if startIdx >= len(packages) {
				break
			}
		}
		ans = min(ans, sum)
	}
	if ans == math.MaxInt64 {
		return -1
	}
	return int(ans % (1000000007))
}

func min(left, right int64) int64 {
	if left < right {
		return left
	}
	return right
}