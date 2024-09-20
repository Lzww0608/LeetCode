type ZigzagIterator struct {
    i int
    q [2][]int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
    i := 0
    q := [2][]int{}
    q[0], q[1] = v1, v2

    return &ZigzagIterator{i, q}
}

func (c *ZigzagIterator) next() int {
    if len(c.q[c.i]) > 0 {
        defer func() { 
            c.q[c.i] = c.q[c.i][1:]
            c.i ^= 1
        }()
        return c.q[c.i][0] 
    }

    c.i ^= 1
    defer func() { 
        c.q[c.i] = c.q[c.i][1:]
    }()
    return c.q[c.i][0] 
}

func (c *ZigzagIterator) hasNext() bool {
	return len(c.q[0]) > 0 || len(c.q[1]) > 0
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */