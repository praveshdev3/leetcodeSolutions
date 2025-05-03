/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slowPointer,fastPointer := head,head
    for fastPointer != nil && fastPointer.Next!=nil {
        slowPointer = slowPointer.Next
        fastPointer = fastPointer.Next.Next
    }
    return slowPointer
}
