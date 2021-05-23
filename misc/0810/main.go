package main

func main() {

}

func xorGame(nums []int) bool {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}
	return ret == 0 || len(nums) % 2 == 0
}