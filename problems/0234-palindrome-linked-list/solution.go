/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }

    mid := findMiddle(head)
    reverse := reverseLinkedList(mid)
    for head!=nil && reverse!=nil{
        if reverse.Val != head.Val{
            return false
        }
        reverse = reverse.Next
        head = head.Next
    }
    return true
}

func findMiddle(head *ListNode) *ListNode{
    slowPointer, fastPointer := head, head
    for fastPointer != nil && fastPointer.Next != nil{
        slowPointer = slowPointer.Next
        fastPointer = fastPointer.Next.Next
    }
    return slowPointer
}

func reverseLinkedList(head *ListNode) *ListNode{
    curr := head
    var prev *ListNode

    for curr!=nil{
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    return prev
}
