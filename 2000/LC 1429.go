type FirstUnique struct {
    a []int
    cnt map[int]int
}


func Constructor(nums []int) FirstUnique {
    cnt := map[int]int{}
    for _, x := range nums {
        cnt[x]++
    }

    return FirstUnique{nums, cnt}
}


func (c *FirstUnique) ShowFirstUnique() int {
    for len(c.a) > 0 {
        if c.cnt[c.a[0]] == 1 {
            return c.a[0]
        }
        c.a = c.a[1:]
    }
    return -1
}


func (c *FirstUnique) Add(value int)  {
    if c.cnt[value] == 0 {
        c.a = append(c.a, value)
    }
    c.cnt[value]++
}


/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */