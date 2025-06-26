type BIT struct {
    f []int
}

func (c *BIT) update(i int) {
    for i < len(c.f) {
        c.f[i] += 1
        i += i & (-i)
    }
}

func (c *BIT) query(i int) (res int) {
    for i > 0 {
        res += c.f[i]
        i -= i & (-i)
    }
    return
}

func kBigIndices(nums []int, k int) (ans int) {
    n := len(nums)
    a := make([]int, n + 2)
    b := make([]int, n + 2)
    pre := &BIT{a}
    suf := &BIT{b}
    left, right := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        left[i] = pre.query(nums[i] - 1)
        right[n-i-1] = suf.query(nums[n-i-1] - 1)
        pre.update(nums[i])
        suf.update(nums[n-i-1])
    }

    for i := 0; i < n; i++ {
        if min(left[i], right[i]) >= k {
            ans++
        }
    }

    return
}