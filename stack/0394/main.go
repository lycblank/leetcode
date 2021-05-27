package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))
}

func decodeString(s string) string {
	stack := make([][]byte, 0, 8)
	numStack := make([]int, 0, 8)
	n := len(s)
	var ans []byte
	for i := 0; i < n; i++{
		if s[i] >= '0' && s[i] <= '9' {
			var num int
			for ; s[i]>='0'&&s[i]<='9'; i++ {
				num = num*10+int(s[i]-'0')
			}
			numStack = append(numStack, num)
		}
		if s[i] == '[' {
			stack = append(stack, ans)
			ans = []byte{}
			continue
		}
		if s[i] == ']' {
			num := numStack[len(numStack)-1]
			ans = bytes.Repeat(ans, num)
			ans = append(stack[len(stack)-1], ans...)
			stack = stack[:len(stack)-1]
			numStack = numStack[:len(numStack)-1]
			continue
		}
		ans = append(ans, s[i])
	}
	for i := len(stack)-1;i>=0;i-- {
		ans = append(stack[i], ans...)
	}
	return string(ans)
}



