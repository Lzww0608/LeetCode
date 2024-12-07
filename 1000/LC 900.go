type RLEIterator struct {
    code []int
}


func Constructor(encoding []int) RLEIterator {
    return RLEIterator{encoding}
}


func (c *RLEIterator) Next(n int) int {
    for len(c.code) > 0 && c.code[0] < n {
        n -= c.code[0]
        c.code = c.code[2:]
    }

    if len(c.code) == 0 {
        return -1
    }
    c.code[0] -= n 
    if c.code[0] == 0 {
        defer func() {c.code = c.code[2:]}()
    }
    return c.code[1]
}


/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(encoding);
 * param_1 := obj.Next(n);
 */