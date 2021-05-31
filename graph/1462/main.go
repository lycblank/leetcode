package main

func main() {

}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	dp := make([][]bool, numCourses)
	for i:=0;i<numCourses;i++{
		dp[i] = make([]bool, numCourses)
		dp[i][i] = true
	}
	for _, pre := range prerequisites {
		dp[pre[0]][pre[1]] = true
	}
	for k := 0; k < numCourses; k++ {
		for i := 0; i < numCourses; i++ {
			for j := 0; j < numCourses; j++ {
				dp[i][j] = dp[i][j] || (dp[i][k] && dp[k][j])
			}
		}
	}
	ans := make([]bool, 0, len(queries))
	for _, q := range queries {
		ans = append(ans, dp[q[0]][q[1]])
	}
	return ans
}
