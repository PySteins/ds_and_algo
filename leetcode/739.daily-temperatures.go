package leetcode

// https://leetcode-cn.com/problems/daily-temperatures/
// 739. 每日温度
import (
	"container/list"
)

func dailyTemperatures(T []int) []int {
	l := len(T)
	if l == 1 {
		return []int{0}
	}
	dayNums := make([]int, l)
	stack := list.New()

	for i := 0; i < l; i++ {
		for stack.Back() != nil && T[stack.Back().Value.(int)] < T[i] {
			top := stack.Back().Value.(int)
			dayNums[top] = i - top
			stack.Remove(stack.Back())
		}
		stack.PushBack(i)
	}
	return dayNums
}
