/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    mid := findMid(head)
    left := sortList(head)
    right := sortList(mid)

    return merge(left,right)
}

func findMid(head *ListNode) *ListNode{
    slow, fast := head, head
    var prev *ListNode
    for fast != nil && fast.Next != nil{
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    prev.Next = nil
    return slow
}

func merge(left, right *ListNode) *ListNode{
    result := &ListNode{}
    head := result

    for left != nil && right != nil{
        if left.Val < right.Val{
            result.Next = left
            left = left.Next
        }else{
            result.Next = right
            right = right.Next
        }
        result = result.Next
    }

    if left != nil{
        result.Next = left
    }else if right != nil{
        result.Next = right
    }

    return head.Next
}
