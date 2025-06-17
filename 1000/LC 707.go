type MyLinkedList struct {
    sz int 
    q []int
}


func Constructor() MyLinkedList {
    return MyLinkedList{0, []int{}}
}


func (c *MyLinkedList) Get(index int) int {
    if index < 0 || index >= c.sz {
        return -1
    }
    return c.q[index]
}


func (c *MyLinkedList) AddAtHead(val int)  {
    c.q = append([]int{val}, c.q...)
    c.sz++
    return
}


func (c *MyLinkedList) AddAtTail(val int)  {
    c.q = append(c.q, val)
    c.sz++
    return
}


func (c *MyLinkedList) AddAtIndex(index int, val int)  {
    if c.sz < index {
        return
    }
    c.q = append(c.q, val)
    c.sz++
    for i := c.sz - 1; i > index; i-- {
        c.q[i] = c.q[i-1]
    }
    c.q[index] = val
    return
}


func (c *MyLinkedList) DeleteAtIndex(index int)  {
    if index >= c.sz {
        return
    }
    for i := index; i < c.sz - 1; i++ {
        c.q[i] = c.q[i + 1]
    }
    c.q = c.q[:c.sz - 1]
    c.sz--
    return
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */