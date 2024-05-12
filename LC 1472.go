type Node struct {
    s string
    next, prev *Node
}
type BrowserHistory struct {
    p *Node
}


func Constructor(homepage string) BrowserHistory {
    p := &Node{homepage, nil, nil}
    return BrowserHistory{p}
}


func (this *BrowserHistory) Visit(url string)  {
    p := &Node{url, nil, this.p}
   
    this.p.next = p 
    this.p = this.p.next

}


func (this *BrowserHistory) Back(steps int) string {
    for i := 0; i < steps && this.p.prev != nil; i++ {
        this.p = this.p.prev
    }
    return this.p.s
}


func (this *BrowserHistory) Forward(steps int) string {
    for i := 0; i < steps && this.p.next != nil; i++ {
        this.p = this.p.next
    }
    return this.p.s
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */