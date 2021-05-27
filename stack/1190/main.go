package main

import "fmt"

func main() {
	s := "(u(love)i)"
	fmt.Println(reverseParentheses(s))
}

func reverseParentheses(s string) string {
	stack := make([][]byte, 0, 8)
	n := len(s)
	var ans []byte
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, ans)
			ans = []byte{}
			continue
		}
		if s[i] == ')' {
			for i := 0; i < len(ans) / 2; i++ {
				ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
			}
			ans = append(stack[len(stack)-1], ans...)
			stack = stack[:len(stack)-1]
			continue
		}
		ans = append(ans, s[i])
	}
	if len(stack) > 0 {
		for i := len(stack)-1;i>=0;i-- {
			ans = append(stack[i], ans...)
		}
	}
	return string(ans)
}