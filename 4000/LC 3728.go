func countStableSubarrays(a []int) int64 {
    n := len(a)
    type pair struct {
        x, sum int
    }
    m := make(map[pair]int)

    ans := 0
    pre := make([]int, n + 1)
    for i := range n {
        pre[i + 1] = pre[i] + a[i]

        if i >= 2 {
            m[pair{a[i - 2], pre[i - 1]}]++
            ans += m[pair{a[i], pre[i + 1] - a[i] * 2}]
        }
    }

    return int64(ans)
}