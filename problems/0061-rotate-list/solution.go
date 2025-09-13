/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    endPointer, startPointer := head, head
    count := 1
    for endPointer.Next != nil{
        count++
        endPointer = endPointer.Next
    }

    k = k%count
    if k == 0 {
        return head
    }
    j := count-k
    for i:=1; i<j; i++{
        startPointer = startPointer.Next
    }
    next := startPointer.Next
    startPointer.Next = nil
    startPointer = next
    endPointer.Next = head
    return startPointer
}
