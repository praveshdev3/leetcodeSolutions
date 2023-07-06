/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
result := &ListNode{}
tmpList := result
for l1 != nil || l2 != nil {
    if l1 != nil {
        tmpList.Val += l1.Val
        l1 = l1.Next
    }
    if l2 != nil {
        tmpList.Val += l2.Val
        l2 = l2.Next
    }
    if tmpList.Val > 9 {
        tmpList.Val -= 10
        tmpList.Next = &ListNode{Val:1}
    } else if l1 != nil || l2 != nil {
        tmpList.Next = &ListNode{}
    }
    tmpList = tmpList.Next

}
return result
}