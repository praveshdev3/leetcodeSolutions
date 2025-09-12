/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil{
        return nil
    }

    dummy1, dummy2 := headA, headB

    for dummy1 != dummy2{
        if dummy1 != nil{
            dummy1 = dummy1.Next
        }else{
            dummy1 = headB
        }

        if dummy2 != nil{
            dummy2 = dummy2.Next
        }else{
            dummy2 = headA
        }
    }

    return dummy1
}
