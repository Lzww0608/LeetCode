type Solution struct {
    prefix []int
}


func Constructor(w []int) Solution {
    n := len(w)
    prefix := make([]int, n + 1)
    for i, x := range w {
        prefix[i+1] = prefix[i] + x
    }

    return Solution{prefix[1:]}
}


func (c *Solution) PickIndex() int {
    x := rand.Intn(c.prefix[len(c.prefix)-1]) + 1

    return sort.SearchInts(c.prefix, x)
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */