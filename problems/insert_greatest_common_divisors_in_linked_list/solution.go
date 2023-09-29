/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getGcd(val1,val2 int) int{
 for val1 != 0{
     val1, val2 = val2 % val1, val1
 }
 return val2
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    res := head
    prev := head
    curr := head.Next
    for curr != nil{
        prev.Next = &ListNode{Val: getGcd(prev.Val, curr.Val),Next: curr}
        prev = curr
        curr = curr.Next
    }
    return res
}