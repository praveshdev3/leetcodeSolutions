/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    prevGroupEnd := dummy

    for {
        kth := getKthNode(prevGroupEnd, k)
        if kth == nil{
            break
        }

        groupStart := prevGroupEnd.Next
        nextGroupStart := kth.Next

        prev := kth.Next
        curr := groupStart

        for curr != nextGroupStart{
            temp := curr.Next
            curr.Next = prev
            prev = curr
            curr = temp
        }

        prevGroupEnd.Next = kth
        prevGroupEnd = groupStart
    }

    return dummy.Next
}

func getKthNode(curr *ListNode, k int) *ListNode{
    for curr != nil && k > 0{
        curr = curr.Next
        k--
    }
    return curr
}
