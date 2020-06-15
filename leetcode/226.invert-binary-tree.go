package leetcode

// https://leetcode-cn.com/problems/invert-binary-tree/
// 226. 翻转二叉树

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
