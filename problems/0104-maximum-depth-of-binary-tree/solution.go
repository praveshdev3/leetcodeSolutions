/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var maximumDepth int
func maxDepth(root *TreeNode) int {
    maximumDepth := 0
    maximumDepth =findDepth(root)
    return maximumDepth
}

func findDepth(node *TreeNode) int{
    if node == nil{
        return 0
    }
    leftDepth := findDepth(node.Left)
    rightDepth := findDepth(node.Right)
    nodalMaxDepth := max(leftDepth,rightDepth)
    maximumDepth = max(maximumDepth,nodalMaxDepth)
    return nodalMaxDepth+1
}
