type MyQueue struct {
    queue []int
}


func Constructor() MyQueue {
    return MyQueue{
        queue: []int{},
    }
}


func (this *MyQueue) Push(x int)  {
    this.queue = append(this.queue, x)
}


func (this *MyQueue) Pop() int {
    if len(this.queue) > 0{
        elem := this.queue[0]
        this.queue = this.queue[1:]
        return elem
    }
    return -1
}


func (this *MyQueue) Peek() int {
    if len(this.queue) > 0{
        return this.queue[0]
    }
    return -1
}


func (this *MyQueue) Empty() bool {
    if len(this.queue) > 0{
        return false
    } 
        return true
    
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
