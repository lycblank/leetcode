package main

func main() {

}

func twoSum(nums []int, target int) []int {
	others := map[int]int{}
	for i, num := range nums {
		if idx, ok := others[target-num]; ok {
			return []int{idx, i}
		}
		others[num] = i
	}
	return []int{-1, -1}
}
