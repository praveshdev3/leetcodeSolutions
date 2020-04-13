class Linklist{
    Node head;
    Node tail;
    Linklist(){
        this.head=null;
        this.tail=null;
    }
}

class Node{
    int data;
    Node next;
    int min;
    Node(int data){
        this.data=data;
        this.next=null;
    }
}

class MinStack {
    Linklist ll;
    public MinStack() {
     ll=new Linklist(); 
    }
    
    public void push(int x) {
       Node nnode=new Node(x);
        if(ll.head==null){
        nnode.min=x;
        ll.head=nnode;
        ll.tail=nnode;
    }
    else
    {
        Node top=ll.head;
        if(x<top.min)
            nnode.min=x;
        else
            nnode.min=top.min;
    }
        nnode.next=ll.head;
        ll.head=nnode;
    }
    
    public void pop() {
        if(ll.head==ll.tail)
        {
            ll.head=ll.tail=null;
            return;
        }
        ll.head=ll.head.next;
    }
    
    public int top() {
        if(ll.head==null)
        {
            return 0;
        }
        return ll.head.data;
    }
    
    public int getMin() {
        if(ll.head==null)
        {
            return 0;
        }
        return ll.head.min;
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */