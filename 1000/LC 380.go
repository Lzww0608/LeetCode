type RandomizedSet struct {
    m map[int]int
    a []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{map[int]int{}, []int{}}
}


func (c *RandomizedSet) Insert(val int) bool {
    if _, ok := c.m[val]; !ok {
        c.a = append(c.a, val)
        c.m[val] = len(c.a) - 1
        return true
    }    
    return false
}


func (c *RandomizedSet) Remove(val int) bool {
    if pos, ok := c.m[val]; ok {
        c.a[pos] = c.a[len(c.a)-1]
        c.m[c.a[pos]] = pos 
        c.a = c.a[:len(c.a)-1]
        delete(c.m, val)
        return true
    }
    return false
}


func (c *RandomizedSet) GetRandom() int {
    pos := rand.Intn(len(c.a)) 
    return c.a[pos]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */