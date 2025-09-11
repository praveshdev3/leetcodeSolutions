/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil || head.Next.Next == nil{
        return head
    }

    evenHead := head.Next
    odd, even := head, head.Next
    
    for even != nil && even.Next != nil{
        odd.Next = odd.Next.Next
        even.Next = even.Next.Next
        odd = odd.Next
        even = even.Next
    }

    odd.Next = evenHead

    return head
}
