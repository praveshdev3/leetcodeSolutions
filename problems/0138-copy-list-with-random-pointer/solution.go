/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    mp := make(map[*Node]*Node)

    endPointer := head

    for endPointer != nil{
        node := &Node{Val: endPointer.Val}
        mp[endPointer] = node
        endPointer = endPointer.Next
    }

    endPointer = head

    for endPointer != nil{
        mp[endPointer].Next = mp[endPointer.Next]
        mp[endPointer].Random = mp[endPointer.Random]
        endPointer = endPointer.Next
    }

    return mp[head]
}
