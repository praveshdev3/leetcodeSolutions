/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil{
        return head
    }

    dummyNode := &ListNode{Next: head}

    slow, fast := dummyNode, head

    for fast != nil && fast.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next
    }

    slow.Next = slow.Next.Next

    return dummyNode.Next
}
