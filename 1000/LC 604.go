type StringIterator struct {
    pre byte
    s []byte
    id, cnt int
}


func Constructor(compressedString string) StringIterator {
    s := []byte(compressedString)
    return StringIterator{' ', s, 0, 0}
}


func (c *StringIterator) Next() byte {
    if c.cnt != 0 {
        c.cnt--
        return c.pre
    } else if c.id >= len(c.s) {
        return ' '
    }
    c.pre = c.s[c.id]
    c.id++
    for c.id < len(c.s) && c.s[c.id] >= '0' && c.s[c.id] <= '9' {
        c.cnt = c.cnt * 10 + int(c.s[c.id] - '0')
        c.id++
    }
    c.cnt--

    return c.pre
}


func (c *StringIterator) HasNext() bool {
    return c.cnt != 0 || c.id < len(c.s)
}


/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */