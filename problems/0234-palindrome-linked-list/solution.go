/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    midNode := getMiddleNode(head)
    reversedNode := getReversedList(midNode)
    return areEquals(head,reversedNode)
}

func getMiddleNode(head *ListNode) *ListNode{
    slowPtr, fastPtr := head,head
    for fastPtr!=nil && fastPtr.Next!=nil{
        slowPtr = slowPtr.Next
        fastPtr = fastPtr.Next.Next
    }
    return slowPtr
}

func getReversedList(head *ListNode) *ListNode{
    var prev *ListNode
    curr := head
    for curr != nil{
        prev, curr, curr.Next = curr, curr.Next, prev
    }
    return prev
}

func areEquals(head, middle *ListNode) bool{
    for head!=nil && middle!=nil{
        if head.Val != middle.Val{
            return false
        }
        middle,head = middle.Next,head.Next
    }
    return true
}


    // arr := make([]int,0)
    // for head != nil{
    //     arr = append(arr, head.Val)
    //     head = head.Next
    // }
    // if len(arr) <= 1{
    //     return true
    // }
    // for i,j := 0,len(arr)-1; i<j; i,j=i+1,j-1{
    //     if arr[i] != arr[j]{
    //         return false
    //     }
    // }
    // return true
