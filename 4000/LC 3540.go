func minTotalTime(forward []int, backward []int, queries []int) int64 {
    n := len(forward)
    pre := make([]int, n + 1)
    for i, x := range forward {
        pre[i + 1] = pre[i] + x
    }

    suf := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        suf[i] = suf[i + 1] + backward[i]
    }

    ans := 0
    i := 0
    for _, q := range queries {
        var cur int
        if i <= q {
            cur = min(pre[q] - pre[i], suf[q + 1] + suf[0] - suf[i + 1])
        } else {
            cur = min(suf[q + 1] - suf[i + 1], pre[q] + pre[n] - pre[i])
        }

        ans += cur
        i = q
        //fmt.Println(cur, ans)
    }

    return int64(ans)
}