/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head==nil || head.Next == nil{
        return false
    }
    slowPointer, fastPointer := head, head
    for fastPointer != nil && fastPointer.Next != nil{
        slowPointer = slowPointer.Next
        fastPointer = fastPointer.Next.Next
        if slowPointer == fastPointer{
            return true
        }
    }
    return false
}
