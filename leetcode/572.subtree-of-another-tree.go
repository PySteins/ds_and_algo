package leetcode

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/subtree-of-another-tree/
// 572. 另一个树的子树

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}
	sb1 := &strings.Builder{}
	sb2 := &strings.Builder{}
	postTraverse(s, sb1)
	postTraverse(t, sb2)
	return strings.Contains(sb1.String(), sb2.String())
}

func postTraverse(node *TreeNode, sb *strings.Builder) {
	if node == nil {
		sb.WriteString("#!")
		return
	}
	postTraverse(node.Left, sb)
	postTraverse(node.Right, sb)
	sb.WriteString(strconv.Itoa(node.Val))
	sb.WriteString("!")
}
