type CustomStack struct {
    st, add []int
    top int
}


func Constructor(maxSize int) CustomStack {
    st := make([]int, maxSize)
    add := make([]int, maxSize)
    return CustomStack{st, add, -1}
}


func (this *CustomStack) Push(x int)  {
    if this.top != len(this.st) - 1 {
        this.top++
        this.st[this.top] = x
    }
}


func (this *CustomStack) Pop() int {
    if this.top == -1 {
        return -1
    }
    res := this.st[this.top] + this.add[this.top]
    if this.top != 0 {
        this.add[this.top-1] += this.add[this.top]
    }
    this.add[this.top] = 0
    this.top--
    return res
}


func (this *CustomStack) Increment(k int, val int)  {
    pos := min(k - 1, this.top)
    if pos >= 0 {
        this.add[pos] += val
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */