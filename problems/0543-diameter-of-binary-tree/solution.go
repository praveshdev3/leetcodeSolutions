/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter = 0
    _ = findDiameter(root)
    return maxDiameter
}

func findDiameter(node *TreeNode) int{
    if node == nil{
        return 0 
    }
    leftDiameter := findDiameter(node.Left)
    rightDiameter := findDiameter(node.Right)
    maxDiameter = max(maxDiameter, leftDiameter+rightDiameter)
    return max(leftDiameter,rightDiameter)+1
}

func max(a,b int) int{
    if a>b{
        return a
    }
    return b
}
