func maximizeTheProfit(n int, offers [][]int) int {
    g := make([][][]int, n)
    for _, v := range offers {
        g[v[0]] = append(g[v[0]], v[1:])
    }

    f := make([]int, n + 2)
    for i := 0; i < n; i++ {
        f[i+1] = max(f[i], f[i+1])
        for _, u := range g[i] {
            a, b := u[0], u[1]
            f[a+1] = max(f[a+1], f[i] + b)
        }
    }

    return slices.Max(f)
}