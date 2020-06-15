package leetcode

// https://leetcode-cn.com/problems/merge-sorted-array/
// 88. 合并两个有序数组
// 类似归并排序的合并
func merge(nums1 []int, m int, nums2 []int, n int) {
	l1, l2, cur := m-1, n-1, len(nums1)-1
	for l1 >= 0 {
		if l2 >= 0 && nums1[l1] > nums2[l2] {
			nums1[cur] = nums1[l1]
			cur--
			l1--
		} else {
			nums1[cur] = nums2[l2]
			cur--
			l2--
		}
	}
}
