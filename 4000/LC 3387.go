func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) (ans float64) {
    solve := func(init string, pairs [][]string, rates []float64) (ans map[string]float64) {
        ans = make(map[string]float64)
        g := map[string][]pair{}
        for i, p := range pairs {
            a, b, c := p[0], p[1], rates[i]
            g[a] = append(g[a], pair{b, c})
            g[b] = append(g[b], pair{a, 1.0 / c})
        }

        var dfs func(string, float64)
        dfs = func(s string, cur float64) {
            ans[s] = cur
            for _, t := range g[s] {
                if ans[t.to] == 0 {
                    dfs(t.to, cur * t.rate)
                }
            }
        }
        dfs(init, 1)
        return 
    }

    a := solve(initialCurrency, pairs1, rates1)
    b := solve(initialCurrency, pairs2, rates2)
    for k, v := range b {
        ans = max(ans, a[k] / v)
    }

    return
}

type pair struct {
    to string
    rate float64
}