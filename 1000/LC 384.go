type Solution struct {
    origin, cur []int
}


func Constructor(nums []int) Solution {
    cur := make([]int, len(nums))
    copy(cur, nums)
    return Solution{
        origin: nums,
        cur: cur,
    }
}


func (c *Solution) Reset() []int {
    copy(c.cur, c.origin)
    return c.cur
}


func (c *Solution) Shuffle() []int {
    n := len(c.cur)

    for i := n - 1; i >= 0; i-- {
        j := rand.Intn(i + 1) 
        c.cur[i], c.cur[j] = c.cur[j], c.cur[i]
    }

    return c.cur
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */