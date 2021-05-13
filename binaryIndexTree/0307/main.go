package main

import "fmt"

func main() {
	nums := []int{1,3,5}
	arr := Constructor(nums)
	fmt.Println(arr.SumRange(0, 2))
	arr.Update(1, 2)
	fmt.Println(arr.SumRange(0, 2))
}

type NumArray struct {
	C []int
	nums []int
}

func Constructor(nums []int) NumArray {
	na := NumArray{
		C:make([]int, len(nums)+1),
		nums:make([]int, len(nums)),
	}
	for i,cnt:=0,len(nums);i<cnt;i++ {
		na.Update(i, nums[i])
	}
	return na
}


func (this *NumArray) Update(index int, val int)  {
	addValue := val - this.nums[index]
	for i := index+1; i < len(this.C); i += this.lowbit(i) {
		this.C[i] += addValue
	}
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Query(right+1) - this.Query(left)
}

func (this *NumArray) Query(x int) int {
	var ans int
	for i := x; i > 0; i -= this.lowbit(i) {
		ans += this.C[i]
	}
	return ans
}

func (this *NumArray) lowbit(x int) int {
	return x&-x
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */