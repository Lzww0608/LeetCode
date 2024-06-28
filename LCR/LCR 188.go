func bestTiming(prices []int) (ans int) {
    if len(prices) == 0 {
        return
    }

    cur := prices[0]
    for _, x := range prices {
        cur = min(cur, x)
        ans = max(ans, x - cur)
    }

    return
}
