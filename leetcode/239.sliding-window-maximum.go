package leetcode

// https://leetcode-cn.com/problems/sliding-window-maximum/
// 239. 滑动窗口最大值

import (
	"container/list"
)

func MaxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	if l == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}
	maxQueue := list.New()
	maxNums := make([]int, l-(k-1))

	for right := 0; right < l; right++ {
		for maxQueue.Front() != nil && nums[right] >= nums[maxQueue.Back().Value.(int)] {
			maxQueue.Remove(maxQueue.Back())
		}
		maxQueue.PushBack(right)

		left := right - (k - 1)
		if left < 0 {
			continue
		}
		if left > maxQueue.Front().Value.(int) {
			maxQueue.Remove(maxQueue.Front())
		}

		maxNums[left] = nums[maxQueue.Front().Value.(int)]

	}
	return maxNums
}
