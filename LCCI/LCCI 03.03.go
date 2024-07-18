type StackOfPlates struct {
    st [][]int
    n int
}


func Constructor(cap int) StackOfPlates {
    st := [][]int{}
    return StackOfPlates{st, cap}
}


func (c *StackOfPlates) Push(val int)  {
    if c.n == 0 {
        return
    }
    if len(c.st) == 0 || len(c.st[len(c.st)-1]) == c.n {
        c.st = append(c.st, []int{})
    }
    c.st[len(c.st)-1] = append(c.st[len(c.st)-1], val)
    return
}


func (c *StackOfPlates) Pop() int {
    return c.PopAt(len(c.st) - 1)
}


func (c *StackOfPlates) PopAt(index int) int {
    ans := -1
    if index >= 0 && index < len(c.st) {
        ans = c.st[index][len(c.st[index])-1]
        c.st[index] = c.st[index][:len(c.st[index])-1]
        if len(c.st[index]) == 0 {
            c.st = append(c.st[:index], c.st[index+1:]...)
        }
    }

    return ans
}


/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */