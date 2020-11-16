/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
      ListNode temp =new ListNode (0);
        ListNode curr=temp;
        while(l1!=null && l2!=null)
        {
            if(l1.val<l2.val)
            {
                curr.next=l1;
                l1=l1.next;
            }else
            {
                curr.next=l2;
                l2=l2.next;
            }
            curr=curr.next;
        }
       while(l1!=null)
       {
           curr.next=l1;
           l1=l1.next;
           curr=curr.next;
       }
        while(l2!=null)
        {
          curr.next=l2;
           l2=l2.next;
           curr=curr.next;  
        }
        return temp.next;
    }
}