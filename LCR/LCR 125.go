type CQueue struct {
    q []int
}


func Constructor() CQueue {
    return CQueue{[]int{}}
}


func (this *CQueue) AppendTail(value int)  {
    this.q = append(this.q, value)
}


func (this *CQueue) DeleteHead() int {
    if len(this.q) == 0 {
        return -1
    }
    res := this.q[0]
    this.q = this.q[1:]
    return res
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
