func houseOfCards(n int) (ans int) {
    f := make([]int, n + 1)
    f[0] = 1

    for i := 1; (i * (3 * i + 1)) / 2 <= n; i++ {
        g := make([]int, n + 1)
        for j := i * 3 - 1; j <= n; j++ {
            if j - 3 * i >= 0 {
                g[j] += g[j - i *  3]
            }
            g[j] += f[j- (3 * i - 1)]
        }
        f = g
        ans += f[n]
    }

    return 
}