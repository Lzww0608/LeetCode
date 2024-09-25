type HitCounter struct {
    a []int
}


func Constructor() HitCounter {
    return HitCounter{[]int{}}
}


func (c *HitCounter) Hit(timestamp int)  {
    c.a = append(c.a, timestamp)
}


func (c *HitCounter) GetHits(timestamp int) (ans int) {
    for i := len(c.a) - 1; i >= 0 && c.a[i] > timestamp - 300; i-- {
        ans++
    }

    return
}


/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */