package main

func main() {

}

func countDigitOne(n int) int {
	dp1 := make([]int, 11)
	tmp := 1
	for i := 1; i <= 10; i, tmp = i + 1, tmp*10{
		dp1[i] = 10*dp1[i-1] + tmp
	}
	bits := make([]int, 0, 10)
	v := n
	for v > 0 {
		bits = append(bits, v % 10)
		v /= 10
	}

	dp2 := make([]int, len(bits)+1)
	if len(bits) > 0 {
		if bits[0] >= 1 {
			dp2[1] = 1
		} else {
			dp2[1] = 0
		}
	}
	tmp = 10
	for i:=2;i<=len(bits);i,tmp=i+1,tmp*10{
		if bits[i-1] == 0 {
			dp2[i] = dp2[i-1]
		} else if bits[i-1] == 1 {
			dp2[i] = dp1[i-1] + dp2[i-1] + (n % tmp) + 1
		} else {
			dp2[i] = bits[i-1]*dp1[i-1] + dp2[i-1] + tmp
		}
	}
	return dp2[len(dp2)-1]
}
