/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return nil
    }

    cycleFound := false

    fastPointer, slowPointer := head, head

    for fastPointer != nil && fastPointer.Next != nil {
        fastPointer = fastPointer.Next.Next
        slowPointer = slowPointer.Next
        if slowPointer == fastPointer{
            cycleFound = true
            break
        }
    }

    if !cycleFound{
        return nil
    }

    fastPointer = head
    for fastPointer != slowPointer{
        fastPointer = fastPointer.Next
        slowPointer = slowPointer.Next
    }

    return slowPointer
}
