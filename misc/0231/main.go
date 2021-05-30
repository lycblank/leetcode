package main

func main() {

}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n > 0 {
		return (n & (n - 1)) == 0
	}
	return false
}
