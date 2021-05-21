package main

func main() {

}

func generateParenthesis(n int) []string {
	arr := make([]byte, 2*n)
	return dfs(arr, 0, 0, 2*n)
}

func dfs(arr []byte, left int, right int, n int) []string {
	if left + right == n {
		if left == right {
			return []string{string(arr)}
		}
		return nil
	}

	arr[left+right] = '('
	ans := dfs(arr, left+1, right, n)
	if left > right {
		arr[left+right] = ')'
		ans = append(ans, dfs(arr, left, right+1, n)...)
	}
	return ans
}

