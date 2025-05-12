/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return helper(root.Left, root.Right)
}

func helper(left, right *TreeNode) bool {
    if left == nil || right == nil {
		return left == right
	}
    if left.Val != right.Val {
        return false
    }
    return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
