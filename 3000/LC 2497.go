func maxStarSum(vals []int, edges [][]int, k int) (ans int) {
    n := len(vals)
    g := make([][]int, n)
    for _, edge := range edges {
        a, b := edge[0], edge[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    ans = math.MinInt32
    for i := 0; i < n; i++ {
        sort.Slice(g[i], func(p, q int) bool {
            return vals[g[i][p]] > vals[g[i][q]]
        })
        sum := vals[i]
        for j := 0; j < min(k, len(g[i])); j++ {
            if vals[g[i][j]] < 0 {
                break
            }
            sum += vals[g[i][j]]
        }
        ans = max(ans, sum)
    }

    return 
}