/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    sum := 0
    ans := &ListNode{}
    curr := ans
    head = head.Next
    for head != nil {
        sum += head.Val
        if head.Val == 0 {
            curr.Next = &ListNode{Val:sum}
            curr = curr.Next
            sum = 0
        }
        head = head.Next
    }
    return ans.Next
}
