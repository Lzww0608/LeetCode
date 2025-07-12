type BIT struct {
    f []int 
}

func (c *BIT) update(i, x int) {
    for ; i < len(c.f); i += i & (-i) {
        c.f[i] = max(c.f[i], x)
    }
}

func (c *BIT) query(i int) (ans int) {
    for ; i > 0; i -= i & (-i) {
        ans = max(ans, c.f[i])
    }

    return
}

func maxProfit(prices []int, profits []int) int {
    n := len(prices)
    mx := slices.Max(prices)
    pre := &BIT{f: make([]int, mx + 1)}
    suf := &BIT{f: make([]int, mx + 1)}
    sufMax := make([]int, n + 1)

    for i := n - 1; i >= 0; i-- {
        x := prices[i]
        sufMax[i] = suf.query(mx - x)
        suf.update(mx - x + 1, profits[i])
    }

    ans := -1
    for i, x := range prices {
        cur := pre.query(x - 1)
        if cur > 0 && sufMax[i] > 0 {
            ans = max(ans, profits[i] + cur + sufMax[i])
        }
        pre.update(x, profits[i])
    }

    return ans
}