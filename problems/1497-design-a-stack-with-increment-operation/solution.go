type CustomStack struct {
    Stack []int
    Top int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{
        Stack: make([]int,maxSize),
        Top: -1,
    }
}


func (this *CustomStack) Push(x int)  {
    if this.Top < len(this.Stack)-1{
        this.Top++
        this.Stack[this.Top] = x
    }
}


func (this *CustomStack) Pop() int {
    if this.Top >= 0{ 
        n := this.Stack[this.Top]
        this.Top--
        return n
    } else{
        return -1
    }
}


func (this *CustomStack) Increment(k int, val int)  {
    if k>this.Top+1{
        k = this.Top+1
    }
    for i:=0; i<k; i++{
        this.Stack[i] = this.Stack[i]+val
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
