/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    var helper func(node *TreeNode) int
    helper = func(node *TreeNode) int{
        if node == nil{
            return 0
        }
        leftMaxSum := max(helper(node.Left),0)
        rightMaxSum := max(helper(node.Right),0)

        maxIfNodeIsRoot := node.Val + leftMaxSum + rightMaxSum
        maxSum = max(maxSum, maxIfNodeIsRoot)

        return node.Val + max(leftMaxSum,rightMaxSum)
    }

    helper(root)
    return maxSum
}

func max(a,b int) int{
    if a>b {
        return a
    }
    return b
}
