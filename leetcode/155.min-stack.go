package leetcode

// https://leetcode-cn.com/problems/min-stack/
// 155. 最小栈

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor1() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, x)
	} else {
		top := this.minStack[len(this.minStack)-1]
		this.minStack = append(this.minStack, min(top, x))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
