type Solution struct {
    cnt map[int][]int
}


func Constructor(nums []int) Solution {
    cnt := make(map[int][]int)
    for i, x := range nums {
        cnt[x] = append(cnt[x], i)
    }

    return Solution{cnt}
}


func (c *Solution) Pick(target int) int {
    a := c.cnt[target]
    n := len(a)
    return a[rand.Intn(n)]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */