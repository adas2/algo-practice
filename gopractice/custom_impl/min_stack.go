package custom

import "math"

// LC # 155
// a stack that supports push, pop, top, and min element in constant time

// define using two stacks
// each stack has both the value and the index as tuple entry
type MinStack struct {
	st1 [][]int //min val stack
	st2 [][]int //regular stack (exclusive)
	cnt int
}

func InitMinStack() MinStack {
	return MinStack{
		st1: make([][]int, 0),
		st2: make([][]int, 0),
		cnt: 0,
	}
}

func (this *MinStack) Push(val int) {
	// choose the appropriate stack
	if len(this.st1) == 0 || val < this.st1[len(this.st1)-1][1] {
		this.st1 = append(this.st1, []int{this.cnt, val})

	} else { // i.e., if val > st1.top()
		this.st2 = append(this.st2, []int{this.cnt, val})
	}

	this.cnt++
}

func (this *MinStack) Pop() {
	// check the indices from the 2 stack tops
	p, q := len(this.st1), len(this.st2)
	if p != 0 && (this.st1[p-1][0] == this.cnt-1) {
		// pop from st1
		this.st1 = this.st1[:p-1]

	} else if q != 0 && (this.st2[q-1][0] == this.cnt-1) {
		// pop from st2
		this.st2 = this.st2[:q-1]

	}

	this.cnt--

}

func (this *MinStack) Top() int {
	p, q := len(this.st1), len(this.st2)
	// top of stack which matches cnt
	if p != 0 && (this.st1[p-1][0] == this.cnt-1) {
		// return st1 top
		return this.st1[p-1][1]

	} else if q != 0 && (this.st2[q-1][0] == this.cnt-1) {
		// return st2 top
		return this.st2[q-1][1]

	}

	// invalid both stack empty
	return math.MinInt32
}

func (this *MinStack) GetMin() int {
	// retuen min of 2 stack tops
	p, q := len(this.st1), len(this.st2)
	minVal := math.MaxInt32

	if p != 0 {
		minVal = min(minVal, this.st1[p-1][1])
	}

	if q != 0 {
		minVal = min(minVal, this.st2[q-1][1])
	}

	return minVal
}
func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

//  logic:
// above approach is space efficient: since we store each element only once is a stack
// another approach is to maintain 2 stack; mainstack --> which has all elemtns
// minStack --> pushed only when val <= minStack.top()
// for pop, we pop mainStack and both only when minStack.top() == mainStack.top()
// GetMin return minStack.top()
