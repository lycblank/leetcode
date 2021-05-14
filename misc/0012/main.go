package main

import "bytes"

func main() {

}

func intToRoman(num int) string {
	var trans = map[int]byte{
		1:'I',
		5: 'V',
		10: 'X',
		50: 'L',
		100: 'C',
		500: 'D',
		1000: 'M',
	}

	var buf bytes.Buffer
	left := num
	x1 := left / 1000
	left = left % 1000
	if x1 > 0 {
		for i:=0;i<x1;i++{
			buf.WriteByte(trans[1000])
		}
	}

	x2 := left / 100
	left = left % 100
	if x2 == 9 {
		buf.WriteByte(trans[100])
		buf.WriteByte(trans[1000])
	} else if x2 == 4 {
		buf.WriteByte(trans[100])
		buf.WriteByte(trans[500])
	} else if x2 >= 5 {
		buf.WriteByte(trans[500])
		for i:=6;i<=x2;i++{
			buf.WriteByte(trans[100])
		}
	} else {
		for i:=1;i<=x2;i++{
			buf.WriteByte(trans[100])
		}
	}


	x3 := left / 10
	left = left % 10
	if x3 == 9 {
		buf.WriteByte(trans[10])
		buf.WriteByte(trans[100])
	} else if x3 == 4 {
		buf.WriteByte(trans[10])
		buf.WriteByte(trans[50])
	} else if x3 >= 5 {
		buf.WriteByte(trans[50])
		for i:=6;i<=x3;i++{
			buf.WriteByte(trans[10])
		}
	} else {
		for i:=1;i<=x3;i++{
			buf.WriteByte(trans[10])
		}
	}

	x4 := left
	if x4 == 9 {
		buf.WriteByte(trans[1])
		buf.WriteByte(trans[10])
	} else if x4 == 4 {
		buf.WriteByte(trans[1])
		buf.WriteByte(trans[5])
	} else if x4 >= 5 {
		buf.WriteByte(trans[5])
		for i:=6;i<=x4;i++{
			buf.WriteByte(trans[1])
		}
	} else {
		for i:=1;i<=x4;i++{
			buf.WriteByte(trans[1])
		}
	}

	return buf.String()
}
