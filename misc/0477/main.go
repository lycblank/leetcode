package main

func main() {

}

func totalHammingDistance(nums []int) int {
	ans, n := 0, len(nums)
	for i := 0; i < 32; i++ {
		var c int
		for _, num := range nums {
			c += (num >> i) & 1
		}
		ans += c * (n-c)
	}
	return ans
}