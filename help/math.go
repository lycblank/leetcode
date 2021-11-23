package help

import "strings"

func max(vals ...int) int {
	ret := vals[0]
	for i:=1;i<len(vals);i++{
		if ret < vals[i] {
			ret = vals[i]
		}
	}
	return ret
}



func isCovered(ranges [][]int, left int, right int) bool {
	for i := left;i<=right;i++{
		isFind := false
		for _, rr := range ranges {
			if i >= rr[0] && i <= rr[1] {
				isFind = true
				break
			}
		}
		if !isFind {
			return false
		}
	}
	return true
}

func chalkReplacer(chalk []int, k int) int {
	sum := 0
	for i:=0;i<len(chalk);i++{
		sum += chalk[i]
	}
	k %= sum
	for i:=0;i<len(chalk);i++{
		if k < chalk[i] {
			return i
		}
		k -= chalk[i]
	}
	return 0
}

func largestMagicSquare(grid [][]int) int {
	strings.ReplaceAll()
	m := len(grid)
	n := len(grid[0])
	sum := make([][]int, m+1)
	for i:=0;i<len(sum);i++ {
		sum[i] = make([]int, n+1)
	}
	for i:=1;i<=m;i++{
		for j:=1;j<=n;j++{
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + grid[i-1][j-1]
		}
	}

	calcLine := func(i,j int, t int) int {
		k := j + t - 1
		return sum[i][k] - sum[i-1][k] - sum[i][j-1] + sum[i-1][j-1]
	}
	calcCol := func(i,j int, t int) int {
		k := i + t - 1
		return sum[k][j] - sum[k][j-1] - sum[i-1][j] + sum[i-1][j-1]
	}
	IsMagicSquare := func(i,j int, t int) bool {
		ans := calcLine(i,j,t)
		for ii := i; ii < i + t; ii++ {
			if ans != calcLine(ii,j,t) {
				 return false
			}
		}

		for jj := j; jj < j + t; jj++ {
			if ans != calcCol(i,jj,t) {
				return false
			}
		}
		var sum1,sum2 int
		for k := 0; k < t; k++ {
			sum1 += grid[i+k][j+k]
			sum2 += grid[i+t-k-1][j+t-1-k]
		}
		if sum1 != ans || sum2 != ans {
			return false
		}
		return true
	}


	k := min(m,n)
	for t:=k;t>=1;t--{
		for i:=1;i<=m-t;i++{
			for j:=1;j<=n-t;j++{
				if IsMagicSquare(i,j,t) {
					return t
				}
			}
		}
	}
	return -1
}

func min(vals ...int) int {
	ret := vals[0]
	for i:=1;i<len(vals);i++{
		if ret > vals[i] {
			ret = vals[i]
		}
	}
	return ret
}