package leetcode

// https://leetcode-cn.com/problems/maximum-binary-tree/
// 654. 最大二叉树

func constructMaximumBinaryTree(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}

	return findRoot(nums, 0, l)
}

func findRoot(nums []int, l, r int) *TreeNode {
	if l == r {
		return nil
	}
	maxIdx := l
	for i := l + 1; i < r; i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}
	return &TreeNode{
		Val:   nums[maxIdx],
		Left:  findRoot(nums, l, maxIdx),
		Right: findRoot(nums, maxIdx+1, r),
	}
}
