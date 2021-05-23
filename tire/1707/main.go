package main

import "sort"

func main() {

}

func maximizeXor(nums []int, queries [][]int) []int {
	anss := make([]int, len(queries))

	sort.Ints(nums)
	for i,cnt := 0,len(queries);i<cnt;i++ {
		queries[i] = append(queries[i], i)
	}
	sort.Slice(queries, func(i,j int) bool {
		return queries[i][1] < queries[j][1]
	})

	tires := &tireTree{}
	var left int
	var n = len(nums)
	for _, q := range queries {
		v, limit, i := q[0], q[1], q[2]
		for left < n && nums[left] <= limit {
			tires.add(nums[left])
			left++
		}
		if left == 0 {
			anss[i] = -1
		} else {
			anss[i] = tires.getMaxXor(v)
		}
	}
	return anss
}


type tireTree struct {
	childern [2]*tireTree
}

func (t *tireTree) add(num int) {
	node := t
	for i:=31;i>=0;i-- {
		v := (num >> i) & 1
		if node.childern[v] == nil {
			node.childern[v] = &tireTree{}
		}
		node = node.childern[v]
	}
}

func (t *tireTree) getMaxXor(val int) int {
	node := t
	var ans int
	for i := 31; i >= 0; i-- {
		v := (val >> i) & 1
		if node.childern[1-v] != nil {
			ans = ans | (1 << i)
			v = 1 - v
		}
		node = node.childern[v]
	}
	return ans
}

