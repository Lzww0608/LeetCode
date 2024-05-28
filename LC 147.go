type MinStack struct {
    st [][]int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{[][]int{}}
}


func (this *MinStack) Push(x int)  {
    if len(this.st) == 0 {
        this.st = append(this.st, []int{x, x})
    } else {
        mn := min(x, this.st[len(this.st)-1][1])
        this.st = append(this.st, []int{x, mn})
    }
}


func (this *MinStack) Pop()  {
    this.st = this.st[:len(this.st)-1]
}


func (this *MinStack) Top() int {
    return this.st[len(this.st)-1][0]
}


func (this *MinStack) GetMin() int {
    return this.st[len(this.st)-1][1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */