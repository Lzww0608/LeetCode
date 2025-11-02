func maxProfit(prices []int, strategy []int, k int) int64 {
    ans := 0
    n := len(prices)

    pre, cur := 0, 0
    mx := 0
    for i := range n {
        ans += prices[i] * strategy[i]

        pre += prices[i] * strategy[i]
        if i < k - 1 {
            if i >= k / 2 {
                cur += prices[i]
            }
        } else {
            cur += prices[i]
            mx = max(mx, cur - pre)
            cur -= prices[i - k / 2 + 1]
            pre -= prices[i - k + 1] * strategy[i - k + 1]
        }
    }


    return int64(ans + mx)
}