func maxScore(n int, k int, stayScore [][]int, travelScore [][]int) int {
    f := make([]int, n + 1)

    for i := 1; i <= k; i++ {
        g := make([]int, n + 1)
        copy(g, f)
        for j := 0; j < n; j++ {
            for p := 0; p < n; p++ {
                if p != j {
                    f[j] = max(f[j], g[p] + travelScore[p][j])
                } else {
                    f[j] = max(f[j], g[p] + stayScore[i - 1][p])
                }
            }
        }
    }

    return slices.Max(f)
}