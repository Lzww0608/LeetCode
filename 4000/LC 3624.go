type BIT struct {
    a []int 
    n int 
}

func (c *BIT) update(i, val int) {
    for ; i < c.n; i += i & (-i) {
        c.a[i] += val
    }
}

func (c *BIT) pre(i int) int {
    res := 0
    for ; i > 0; i -= i & (-i) {
        res += c.a[i]
    }

    return res
}

func (c *BIT) query(l, r int) int {
    return c.pre(r) - c.pre(l - 1)
}

func calc(x int) int {
    res := 0
    for x > 1 {
        res++
        x = bits.OnesCount(uint(x))
    }

    return res
}

func popcountDepth(nums []int64, queries [][]int64) []int {
    f := [6]BIT{}
    n := len(nums)
    for i := range f {
        a := make([]int, n + 1)
        f[i] = BIT{a, n + 1}
    }

    update := func(i, val int) {
        x := calc(int(nums[i]))
        f[x].update(i + 1, val)
    }

    ans := []int{}
    for i := range n {
        update(i, 1)
    }

    for _, v := range queries {
        if v[0] == 1 {
            ans = append(ans, f[int(v[3])].query(int(v[1]) + 1, int(v[2]) + 1))
        } else {
            i, t := int(v[1]), v[2]
            x := calc(int(nums[i]))
            f[x].update(i + 1, -1)
            nums[i] = t
            x = calc(int(nums[i]))
            f[x].update(i + 1, 1)
        }
    }

    return ans
}