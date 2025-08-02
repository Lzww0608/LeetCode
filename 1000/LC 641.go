type MyCircularDeque struct {
    l, r []int
    cap int
}


func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{[]int{}, []int{}, k}
}


func (c *MyCircularDeque) InsertFront(value int) bool {
    if c.IsFull() {
        return false
    }
    c.l = append(c.l, value)
    return true
}


func (c *MyCircularDeque) InsertLast(value int) bool {
    if c.IsFull() {
        return false
    }
    c.r = append(c.r, value)
    return true
}


func (c *MyCircularDeque) DeleteFront() bool {
    if c.IsEmpty() {
        return false
    }
    if len(c.l) > 0 {
        c.l = c.l[:len(c.l)-1]
    } else {
        c.r = c.r[1:]
    }
    return true
}


func (c *MyCircularDeque) DeleteLast() bool {
    if c.IsEmpty() {
        return false
    }
    if len(c.r) > 0 {
        c.r = c.r[:len(c.r)-1]
    } else {
        c.l = c.l[1:]
    }
    return true
}


func (c *MyCircularDeque) GetFront() int {
    if c.IsEmpty() {
        return -1
    }
    if len(c.l) > 0 {
        return c.l[len(c.l)-1]
    }
    return c.r[0]
}


func (c *MyCircularDeque) GetRear() int {
    if c.IsEmpty() {
        return -1
    }
    if len(c.r) > 0 {
        return c.r[len(c.r)-1]
    }
    return c.l[0]
}


func (c *MyCircularDeque) IsEmpty() bool {
    return len(c.l) + len(c.r) == 0
}


func (c *MyCircularDeque) IsFull() bool {
    return len(c.l) + len(c.r) == c.cap
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */