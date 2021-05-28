package main

func main() {

}

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i,cnt:=0,len(dp);i<cnt;i++{
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0]=true

	match := func(i,j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	for i,lens := 0,len(s);i<=lens;i++{
		for j,lenp := 1,len(p); j <=lenp; j++ {
			if p[j-1] != '*' {
				if match(i, j) {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				dp[i][j] = dp[i][j-2]
				if match(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
