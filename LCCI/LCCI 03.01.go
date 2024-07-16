type TripleInOne struct {
    st []int
    top []int
}


func Constructor(stackSize int) TripleInOne {
    st := make([]int, stackSize * 3)
    top := make([]int, 3)
    for i := range top {
        top[i] = i
    }

    return TripleInOne{st, top}
}


func (c *TripleInOne) Push(stackNum int, value int)  {
    if c.top[stackNum] >= len(c.st) {
        return
    }
    c.st[c.top[stackNum]] = value
    c.top[stackNum] += 3
}


func (c *TripleInOne) Pop(stackNum int) int {
    if c.IsEmpty(stackNum) {
        return -1
    }
    c.top[stackNum] -= 3
    return c.st[c.top[stackNum]]
}


func (c *TripleInOne) Peek(stackNum int) int {
    if c.IsEmpty(stackNum) {
        return -1
    }
    ans := c.st[c.top[stackNum]-3]
    return ans
}


func (c *TripleInOne) IsEmpty(stackNum int) bool {
    return c.top[stackNum] - 3 < 0 
}


/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */