/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil{
        return false
    }
    
    fastPointer, slowPointer := head, head
    for fastPointer != nil && fastPointer.Next != nil{
        fastPointer = fastPointer.Next.Next
        slowPointer = slowPointer.Next
        if slowPointer == fastPointer{
            return true
        }
    }

    return false
}
