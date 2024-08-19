type CombinationIterator struct {
    s []string
    cur int
}


func Constructor(characters string, combinationLength int) CombinationIterator {
    arr := []string{}
    n := len(characters)
    for i := 1; i < (1 << n); i++ {
        if bits.OnesCount(uint(i)) == combinationLength {
            tmp := []byte{}
            for j := 0; j < n; j++ {
                if i & (1 << j) != 0 {
                    tmp = append(tmp, characters[j])
                }
            }
            arr = append(arr, string(tmp))
        }
    }
    sort.Strings(arr)
    return CombinationIterator{arr, 0}
}


func (c *CombinationIterator) Next() string {
    ans := c.s[c.cur]
    
    c.cur++
    return ans
}


func (c *CombinationIterator) HasNext() bool {
    return c.cur < len(c.s)
}


/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */