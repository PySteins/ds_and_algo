package leetcode

import "math"

// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// 104. 二叉树的最大深度

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)+1), float64(maxDepth(root.Right)+1)))
}
