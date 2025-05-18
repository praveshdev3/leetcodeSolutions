/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil{
        return [][]int{}
    }
    queue := []*TreeNode{root}
    res := [][]int{}

    for length := 0; len(queue) > 0; length++{
        res = append(res, []int{})
        for lengthSize := len(queue); lengthSize > 0; lengthSize-- {
            if queue[0].Left != nil{
                queue = append(queue,queue[0].Left)
            }
            if queue[0].Right != nil{
                queue = append(queue,queue[0].Right)
            }
            res[length] = append(res[length],queue[0].Val)
            queue = queue[1:]
        }
    }

    return res
}
