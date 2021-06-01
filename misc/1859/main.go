package main

import (
	"strings"
)

func main() {

}

func sortSentence(s string) string {
	ss := strings.Split(s, " ")
	ans := make([]string, len(ss))
	for i,cnt:=0,len(ss);i<cnt;i++ {
		idx := int(ss[i][len(ss[i])-1] - '0') - 1
		ans[idx] = ss[i][:len(ss[i])-1]
	}
	return strings.Join(ans, " ")
}

