package main

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	sum := make([][]int, len(matrix)+1)
	for i:=0;i<len(sum);i++{
		sum[i] = make([]int, len(matrix[0])+1)
	}
	for i:=1;i<=len(matrix);i++{
		for j:=1;j<=len(matrix[i-1]);j++{
			sum[i][j] = sum[i-1][j] + sum[i][j-1]-sum[i-1][j-1]+matrix[i-1][j-1]
		}
	}

	var ans int
	for i:=0;i<=len(matrix);i++{
		for j:=0;j<=len(matrix[0]);j++{
			for p:=1;p<=i;p++{
				for q:=1;q<=j;q++{
					if sum[i][j]-sum[p-1][j]-sum[i][q-1]+sum[p-1][q-1] == target {
						ans++
					}
				}
			}
		}
	}
	return ans
}


