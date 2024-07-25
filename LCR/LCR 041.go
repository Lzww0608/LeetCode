type MovingAverage struct {
    a []int
    sum, size int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    a := []int{}
    return MovingAverage{a, 0, size}
}


func (c *MovingAverage) Next(val int) float64 {
    c.a = append(c.a, val)
	c.sum += val
	if len(c.a) > c.size {
		c.sum -= c.a[0]
		c.a = c.a[1:]
	}
	return float64(c.sum) / float64(min(len(c.a), c.size))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */