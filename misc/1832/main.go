package main

func main() {

}

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	var ans int
	for i,cnt:=0,len(sentence);i<cnt;i++{
		ans |= 1 << int(sentence[i]-'a')
	}
	return ans == 0x03ffffff
}

