type LRUCache struct {
    Head, Tail *Node
    Mp map[int]*Node
    Capacity int
}

type Node struct{
    Key,Value int
    Prev,Next *Node
}

func newNode(key,value int) *Node{
    return &Node{
        Key: key,
        Value: value,
    }
}

func Constructor(capacity int) LRUCache {
    head := newNode(0,0)
    tail := newNode(0,0)
    head.Next = tail
    tail.Prev = head
    return LRUCache{
        Head: head,
        Tail: tail,
        Mp: make(map[int]*Node),
        Capacity: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    if n,ok:=this.Mp[key]; ok{
        this.remove(n)
        this.insert(n)
        return n.Value
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if _,ok := this.Mp[key];ok{
        this.remove(this.Mp[key])
    }
    if len(this.Mp) == this.Capacity{
        this.remove(this.Tail.Prev)
    }
    this.insert(newNode(key,value))
}

func (this *LRUCache) remove(node *Node){
    delete(this.Mp,node.Key)
    node.Prev.Next=node.Next
    node.Next.Prev=node.Prev
}

func (this *LRUCache) insert(node *Node){
    this.Mp[node.Key] = node
    next := this.Head.Next
    this.Head.Next = node
    node.Prev = this.Head
    next.Prev = node
    node.Next = next
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
