/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    leftLeg := leftLegLen(root)
    rightLeg := rightLegLen(root)
    if leftLeg == rightLeg{
        return int(math.Pow(2.0,float64(leftLeg)) - 1)
    } else {
        return 1 + countNodes(root.Left) + countNodes(root.Right)
    }
}

func leftLegLen(root *TreeNode) int{
    if root == nil{
        return 0
    }
    return 1 + leftLegLen(root.Left)
}

func rightLegLen(root *TreeNode) int{
    if root == nil{
        return 0
    }
    return 1 + rightLegLen(root.Right)
}
    // if root != nil {
    //     return 1+countNodes(root.Left)+countNodes(root.Right)
    // }
    // return 0
