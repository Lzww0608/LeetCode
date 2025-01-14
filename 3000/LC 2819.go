func minimumRelativeLosses(prices []int, queries [][]int) []int64 {
    n := len(prices)
    sort.Ints(prices)
    pre := make([]int, n + 1)
    for i, x := range prices {
        pre[i+1] = pre[i] + x
    }

    ans := make([]int64, len(queries))
    for i, v := range queries {
        k, m := v[0], v[1]
        if m == n {
            pos := sort.SearchInts(prices, k + 1)
            ans[i] = int64(pre[pos] + k * (n - pos) * 2 - (pre[n] - pre[pos]))
            continue
        }

        l, r := 0, m
        pos := m
        for l < r {
            mid := l + ((r - l) >> 1)
            start, end := mid, mid + n - m 
            if prices[end] - k < k - prices[start] {
                l = mid + 1
            } else {
                pos = mid
                r = mid
            }
        }

        ans[i] = int64(pre[pos] + 2 * k * (n - pos - (n - m)) - (pre[n] - pre[pos+n-m]))
    }

    return ans
}
