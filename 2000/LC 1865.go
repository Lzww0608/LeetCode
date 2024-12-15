type FindSumPairs struct {
    a, b []int 
    cnt map[int]int
}


func Constructor(nums1 []int, nums2 []int) FindSumPairs {
    cnt := map[int]int{}
    for _, x := range nums2 {
        cnt[x]++
    }
    return FindSumPairs{nums1, nums2, cnt}
}


func (c *FindSumPairs) Add(index int, val int)  {
    c.cnt[c.b[index]]--
    c.b[index] += val
    c.cnt[c.b[index]]++
}


func (c *FindSumPairs) Count(tot int) (ans int) {
    for _, x := range c.a {
        ans += c.cnt[tot-x]
    }
    return
}


/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */