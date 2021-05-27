package main

func main() {

}

func hammingDistance(x int, y int) int {
	ans,num:=0,x^y
	for num > 0 {
		ans, num = ans+1, num & (num-1)
	}
	return ans
}
