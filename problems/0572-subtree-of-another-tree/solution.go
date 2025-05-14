/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
      if root == nil{
        return false
      }
      if isSameTree(root,subRoot){
        return true
      } 

      return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p, q *TreeNode) bool {
    if p == nil || q == nil{
        return p == q
    }

    if p.Val != q.Val{
        return false
    }

   return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
