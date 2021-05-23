package main

func main() {

}

func findMaximumXOR(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	tires := &tireTree{}
	var ans int
	for i, num := range nums {
		if i > 0 {
			v := tires.getMaxXor(num)
			if v > ans {
				ans = v
			}
		}
		tires.add(num)
	}
	return ans
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
