package main

import "fmt"

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(topKFrequent(words, 2))
}

func topKFrequent(words []string, k int) []string {
	wm := make(map[string]int)
	for i,cnt:=0,len(words);i<cnt;i++{
		wm[words[i]]++
	}
	ws := make([]*WordInfo, 0, len(wm))
	for k,v:=range wm {
		ws = append(ws, &WordInfo{Word: k,Count: v})
	}
	return topK(ws, k)
}

func topK(words []*WordInfo, k int) []string {
	if len(words) <= 0 {
		return nil
	}
	if len(words) == 1 {
		return []string{words[0].Word}
	}
	right := len(words)-1
	left, right := 0, len(words) - 1
	for left < right {
		for  Less(words[right], words[left]) && left < len(words) {
			left++
			continue
		}
		words[left], words[right] = words[right], words[left]
		for Less(words[right], words[left]) && right >= 0 {
			right--
			continue
		}
		words[left], words[right] = words[right], words[left]
	}
	if k <= right { // 全部在左边
		return topK(words[:right], k)
	}
	// 一部分在左边，一部分在右边
	return append(topK(words[:right], right), topK(words[right:], k - right)...)
}

func Less(left,right *WordInfo) bool {
	if left.Count < right.Count {
		return true
	}
	if left.Count == right.Count {
		return left.Word > right.Word
	}
	return false
}

type WordInfo struct {
	Word string
	Count int
}