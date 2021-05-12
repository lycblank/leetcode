package main

func main() {

}

func decode(encoded []int) []int {
	arr := make([]int, len(encoded)+1)

	sumXOR := 0
	for i,cnt:=1,len(arr);i<=cnt;i++{
		sumXOR = sumXOR ^ i
	}

	tmp := 0
	for i,cnt:=0,len(encoded); i < cnt; i+=2 {
		tmp = tmp ^ encoded[i]
	}

	arr[len(arr)-1] = sumXOR ^ tmp

	for i := len(arr) -1; i>=1;i--{
		arr[i-1] = arr[i] ^ encoded[i-1]
	}
	return arr
}