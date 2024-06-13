type MinStack struct {
    st [][]int
}


func Constructor() MinStack {
    return MinStack{[][]int{}}
}


func (this *MinStack) Push(val int)  {
    if len(this.st) == 0 {
        this.st = append(this.st, []int{val, val})
        return
    }
    mn := min(val, this.st[len(this.st)-1][1])
    this.st = append(this.st, []int{val, mn})
    return 
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
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
