func maxTaxiEarnings(n int, rides [][]int) int64 {
    f := make([]int, n + 1)
    sort.Slice(rides, func(i, j int) bool {
        return rides[i][0] < rides[j][0]
    })

    m := len(rides)
    for i, j := 1, 0; i <= n; i++ {
        f[i] = max(f[i], f[i - 1])
        for j < m && rides[j][0] == i {
            r, x := rides[j][1], rides[j][2]
            f[r] = max(f[r], f[i] + r - i + x)
            j++
        }
    }

    return int64(f[n])
}