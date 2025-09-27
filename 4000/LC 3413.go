func maximumCoins(coins [][]int, k int) int64 {
    n := len(coins)
    ans := solve(coins, k)

    for i := range n {
        coins[i][0], coins[i][1] = -coins[i][1], -coins[i][0]
    }

    return max(ans, solve(coins, k))
}

func solve(coins [][]int, k int) int64 {
    n := len(coins)
    sort.Slice(coins, func(i, j int) bool {
        return coins[i][0] < coins[j][0]
    })

    pre := make([]int, n + 1)
    for i, v := range coins {
        pre[i + 1] = pre[i] + (v[1] - v[0] + 1) * v[2]
    }

    ans := 0
    for i, v := range coins {
        l := v[0]
        j := sort.Search(n, func(j int) bool {
            return coins[j][1] >= l + k - 1
        })
        if j == n {
            ans = max(ans, pre[n] - pre[i])
            break
        }

        add := coins[j][1] - l + 1 - k
        a := min(coins[j][1] - coins[j][0] + 1, add) * coins[j][2]
        ans = max(ans, pre[j + 1] - pre[i] - a)
    }

    return int64(ans)
} 