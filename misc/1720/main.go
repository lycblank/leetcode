package main

func main() {

}

func decode(encoded []int, first int) []int {
	arr := make([]int, len(encoded)+1)
	arr[0] = first

	for i,cnt:=1,len(arr); i < cnt;i++ {
		arr[i] = arr[i-1]^encoded[i-1]
	}
	return arr
}







