package custom

import (
	"math/rand"
)

// LC # 380

type RandomizedSet struct {
	set map[int]int // key: input_number, value: index of arr
	arr []int
	cnt int //optimizes time compared to len() operation
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		set: make(map[int]int),
		arr: make([]int, 0),
		cnt: 0,
	}
}

// each function works in average O(1) time complexity.
func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.set[val]; exists {
		return false
	}
	this.set[val] = this.cnt
	this.arr = append(this.arr, val)
	this.cnt++
	// fmt.Printf("Set with len %d: %v\nArr with len %d: %v\n", len(this.set), this.set, this.cnt, this.arr)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.set[val]; !exists {
		return false
	}
	idx := this.set[val]
	// swap the last element in arr with the deleted item index
	lastElem := this.arr[this.cnt-1]
	this.arr[idx] = lastElem
	this.arr = this.arr[:this.cnt-1]
	this.cnt--
	// tricky: adjust index for swapped element
	this.set[lastElem] = idx
	delete(this.set, val)
	// fmt.Printf("Set with len %d: %v\nArr with len %d: %v\n", len(this.set), this.set, this.cnt, this.arr)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	// generate random index based on len of map
	// rand.Seed(time.Now().UnixNano())
	// fmt.Printf("Set with len %d: %v\nArr with len %d: %v\n", len(this.set), this.set, this.cnt, this.arr)
	idx := rand.Intn(this.cnt)

	return this.arr[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
