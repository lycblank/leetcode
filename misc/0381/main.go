package main

import "math/rand"

func main() {

}


type RandomizedCollection struct {
	kv map[int]map[int]struct{}
	nums []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	rc := RandomizedCollection{
		kv:map[int]map[int]struct{}{},
	}
	return rc
}


/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	indexs, ok := this.kv[val]
	if !ok {
		indexs = map[int]struct{}{}
		this.kv[val] = indexs
	}

	indexs[len(this.nums)] = struct{}{}
	this.nums = append(this.nums, val)
	return !ok
}


/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	indexs, ok := this.kv[val]
	if !ok {
		return ok
	}

	var idx int
	for idx = range indexs {
		break
	}
	n := len(this.nums)
	this.nums[idx] = this.nums[n-1]
	delete(indexs, idx)
	delete(this.kv[this.nums[idx]], n-1)
	if idx < n - 1 {
		this.kv[this.nums[idx]][idx] = struct{}{}
	}
	if len(indexs) <= 0 {
		delete(this.kv, val)
	}
	this.nums = this.nums[:n-1]
	return true
}


/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	if len(this.nums) <= 0 {
		return -1
	}
	return this.nums[rand.Intn(len(this.nums))]
}







/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */