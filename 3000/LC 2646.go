func minimumTotalPrice(n int, edges [][]int, price []int, trips [][]int) int {
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    cnt := make([]int, n)
    for _, v := range trips {
        a, b := v[0], v[1]
        
        var dfs func(int, int) bool
        dfs = func(root, fa int) bool {
            cnt[root]++
            if root == b {
                return true
            }
            for _, x := range g[root] {
                if x != fa {
                    if dfs(x, root) {
                        return true
                    }
                }
            }

            cnt[root]--
            return false
        }
        dfs(a, -1)

    }

    var f func(int, int) (int, int)
    f = func(root, fa int) (notChosen, chosen int) {
        notChosen = cnt[root] * price[root]
        chosen = cnt[root] * price[root] / 2
        for _, x := range g[root] {
            if x != fa {
                a, b := f(x, root)
                notChosen += min(a, b)
                chosen += a
            }
        }
        return 
    }
    nc, c := f(0, -1)

    return min(nc, c)
}
