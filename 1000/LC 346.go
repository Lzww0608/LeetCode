type MovingAverage struct {
    a []int
    sum, n int
}


func Constructor(size int) MovingAverage {
    sum, n := 0, size
    a := []int{}
    return MovingAverage{a, sum, n}
}


func (c *MovingAverage) Next(val int) float64 {
    c.a = append(c.a, val)
    c.sum += val 
    if len(c.a) <= c.n {
        return float64(c.sum) / float64(len(c.a))
    }
    c.sum -= c.a[0]
    c.a = c.a[1:]
    return float64(c.sum) / float64(c.n)
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */