/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    accumulation := 0 

    var helper func (node *TreeNode) *TreeNode
    helper = func(node *TreeNode) *TreeNode{
        if node != nil{
            helper(node.Right)

            accumulation += node.Val
            node.Val = accumulation

            helper(node.Left)
        }
        return node
    }

    return helper(root)
}
