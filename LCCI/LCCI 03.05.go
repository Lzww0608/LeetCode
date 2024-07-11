type SortedStack struct {
    stMin []int
    stMax []int
}


func Constructor() SortedStack {
    stMax, stMin := []int{}, []int{}
    return SortedStack{stMin, stMax}
}


func (c *SortedStack) Push(val int)  {
    for len(c.stMax) > 0 && val < c.stMax[len(c.stMax)-1] {
        t := c.stMax[len(c.stMax)-1]
        c.stMax = c.stMax[:len(c.stMax)-1]
        c.stMin = append(c.stMin, t)
    }

    for len(c.stMin) > 0 && val > c.stMin[len(c.stMin)-1] {
        t := c.stMin[len(c.stMin)-1]
        c.stMin = c.stMin[:len(c.stMin)-1]
        c.stMax = append(c.stMax, t)
    }
    
    c.stMin = append(c.stMin, val)
}


func (c *SortedStack) Pop()  {
    for len(c.stMax) > 0 {
        t := c.stMax[len(c.stMax)-1]
        c.stMax = c.stMax[:len(c.stMax)-1]
        c.stMin = append(c.stMin, t)
    }
    if !c.IsEmpty() {
        c.stMin = c.stMin[:len(c.stMin)-1]
    }
    
}


func (c *SortedStack) Peek() int {
    if c.IsEmpty() {
        return -1
    }
    for len(c.stMax) > 0 {
        t := c.stMax[len(c.stMax)-1]
        c.stMax = c.stMax[:len(c.stMax)-1]
        c.stMin = append(c.stMin, t)
    }
    return c.stMin[len(c.stMin)-1]
}


func (c *SortedStack) IsEmpty() bool {
    return len(c.stMax) == 0 && len(c.stMin) == 0
}


/**
 * Your SortedStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.IsEmpty();
 */