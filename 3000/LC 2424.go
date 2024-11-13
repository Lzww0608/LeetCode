type LUPrefix struct {
    p int
    a []bool
}


func Constructor(n int) LUPrefix {
    a := make([]bool, n + 2)
    return LUPrefix{1, a}
}


func (c *LUPrefix) Upload(video int)  {
    c.a[video] = true
}


func (c *LUPrefix) Longest() int {
    for c.a[c.p] {
        c.p++
    }
    return c.p - 1
}


/**
 * Your LUPrefix object will be instantiated and called as such:
 * obj := Constructor(n);
 * obj.Upload(video);
 * param_2 := obj.Longest();
 */