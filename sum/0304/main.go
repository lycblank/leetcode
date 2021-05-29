package main

func main() {

}

type NumMatrix struct {
	sum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix)+1)
	for i := 0; i < len(sum); i++ {
		sum[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[i-1]); j++ {
			sum[i][j] = sum[i-1][j] + sum[i][j-1] - sum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{
		sum: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1++
	col1++
	row2++
	col2++
	return this.sum[row2][col2] - this.sum[row1-1][col2] -
		this.sum[row2][col1-1] + this.sum[row1-1][col1-1]
}

