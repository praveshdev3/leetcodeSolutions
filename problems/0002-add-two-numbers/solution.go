/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil || l2 == nil{
        return nil
    }

    result := &ListNode{}
    res := result

    carry := 0
    for l1 != nil || l2 != nil || carry > 0{
        node := &ListNode{}
        if l1 != nil{
            carry += l1.Val
            l1 = l1.Next
        }
        if l2 != nil{
            carry += l2.Val
            l2 = l2.Next
        }
        node.Val = carry%10
        carry = carry/10 
        result.Next = node
        result = result.Next
    }

    return res.Next
}
