/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    res,_ := dfs(root)
    return res
}

func dfs(root *TreeNode) (bool,int){
    if root == nil{
        return true,0
    }
    isLeftSubtreeBalanced, leftSubtreeHeight := dfs(root.Left)
    isRightSubtreeBalanced, rightSubtreeHeight := dfs(root.Right)
    diff := abs(leftSubtreeHeight - rightSubtreeHeight)
    if isLeftSubtreeBalanced && isRightSubtreeBalanced && diff <= 1{
        return true, 1+max(leftSubtreeHeight,rightSubtreeHeight)
    }
    return false,-1
}

func abs(num int)int{
    if num<0{
        return -num
    }
    return num
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
