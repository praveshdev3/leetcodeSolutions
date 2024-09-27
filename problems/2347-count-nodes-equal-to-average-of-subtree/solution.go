/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
    result := 0
    _,_ = traverse(root, &result)
    return result
}

func traverse(root *TreeNode, result *int) (int,int){
    if root == nil {
        return 0,0
    }
    leftSum, leftCount := traverse(root.Left,result)
    rightSum, rightCount := traverse(root.Right,result)
    sum := leftSum+rightSum+root.Val
    count := 1+leftCount+rightCount
    average := sum/count
    if average == root.Val{
        *result++
    }
    return sum,count
}
