type Bitset struct {
    a []byte
    n int
    f bool
}


func Constructor(size int) Bitset {
    a := make([]byte, size)
    for i := range a {
        a[i] = '0'
    }
    return Bitset{a, 0, false}
}


func (c *Bitset) Fix(idx int)  {
    if !c.f {
        if c.a[idx] == '0' {
            c.n++
        }
        c.a[idx] = '1'
    } else {
        if c.a[idx] == '1' {
            c.n++
        }
        c.a[idx] = '0'
    }
    
}


func (c *Bitset) Unfix(idx int)  {
    if !c.f {
        if c.a[idx] == '1' {
            c.n--
        }
        c.a[idx] = '0'
    } else {
        if c.a[idx] == '0' {
            c.n--
        }
        c.a[idx] = '1'
    }
}


func (c *Bitset) Flip()  {
    c.f = !c.f
    c.n = len(c.a) - c.n
}


func (c *Bitset) All() bool {
    return len(c.a) == c.n
}


func (c *Bitset) One() bool {
    return c.n > 0
}


func (c *Bitset) Count() int {
    return c.n
}


func (c *Bitset) ToString() string {
    result := make([]byte, len(c.a))
	if c.f {
		for i := range c.a {
			if c.a[i] == '0' {
				result[i] = '1'
			} else {
				result[i] = '0'
			}
		}
	} else {
		copy(result, c.a)
	}
	return string(result)
}


/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */