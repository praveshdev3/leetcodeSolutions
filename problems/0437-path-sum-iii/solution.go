/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
curPath := make([]int,0)
count := 0
recursiveSum(root,targetSum,curPath,&count)
return count
}

func recursiveSum(root *TreeNode, targetSum int, curPath []int, count *int){
    if root == nil{
        return
    }
    curPath = append(curPath, root.Val)
    sum := 0
    for i := len(curPath) - 1; i>=0; i--{
        sum = sum + curPath[i]
        if sum == targetSum {
            *count++
        }
    }
    recursiveSum(root.Left,targetSum,curPath,count)
    recursiveSum(root.Right,targetSum,curPath,count)
    return
}
    // if root == nil {
    //     return 0
    // }
    // return sumUp(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
// func sumUp(root *TreeNode, target int) int{
//     if root == nil{
//         return 0
//     }
//     current := target - root.Val
//     c := 0
//     if current == 0{
//         c = 1
//     }
//     return c + sumUp(root.Left, current) + sumUp(root.Right, current)
// }
