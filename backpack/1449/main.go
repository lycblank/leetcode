package main

func main() {

}

func largestNumber(cost []int, target int) string {
	dp := make([]string, target+1)
	for i:=1;i<=len(cost);i++{
		if cost[i-1] <= target {
			dp[cost[i-1]] = string([]byte{byte(i + '0')})
		}
	}

	for i:=1;i<=len(cost);i++{
		v := cost[i-1]
		for j:=v;j<=target;j++{
			if len(dp[j-v]) <= 0 {
				continue
			}
			dp[j] = max(dp[j], dp[j-v] + string([]byte{byte(i+'0')}))
		}
	}
	if len(dp[target]) <= 0 {
		return "0"
	}
	ans := []byte(dp[target])
	for i,cnt:=0,len(ans);i<cnt/2;i++{
		ans[i], ans[cnt-1-i] = ans[cnt-1-i], ans[i]
	}
	return string(ans)
}

func max(left string, right string) string {
	if len(left) > len(right) {
		return left
	}
	if len(right) > len(left) {
		return right
	}

	for i := len(left)-1;i>=0;i--{
		if left[i] > right[i] {
			return left
		}
		if right[i] > left[i] {
			return right
		}
	}
	return left
}