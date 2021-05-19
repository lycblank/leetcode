package main

import "sort"

func main() {

}

func kthLargestValue(matrix [][]int, k int) int {
	rets := make([]int, 0, len(matrix)*len(matrix[0]))
	results := make([][]int, len(matrix))
	for i,cnt:=0,len(results);i<cnt;i++{
		results[i] = make([]int, len(matrix[i]))
	}
	for i,cnt:=0,len(matrix);i<cnt;i++{
		for j,cnt2:=0,len(matrix[i]);j<cnt2;j++{
			results[i][j] = matrix[i][j]
			if i - 1 >= 0 {
				results[i][j] ^= results[i-1][j]
			}
			if j - 1 >= 0 {
				results[i][j] ^= results[i][j-1]
			}
			if i - 1 >= 0 && j - 1 >= 0 {
				results[i][j] ^= results[i-1][j-1]
			}
			rets = append(rets, results[i][j])
		}
	}
	sort.Ints(rets)
	return rets[len(rets)-k]
}
