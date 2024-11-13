func countGoodNodes(edges [][]int) (ans int) {
    n := len(edges) + 1
    g := make([][]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    var dfs func(int, int) int
    dfs = func(x, fa int) (res int) {
        cnt := 0
        for _, y := range g[x] {
            if y == fa {
                continue
            }
            t := dfs(y, x)
            if cnt != 0 && cnt != t {
                cnt = -1
            } else {
                cnt = t 
            }
            res += t 
        }

        if cnt != -1 {
            ans++
        }

        return res + 1
    }
    dfs(0, -1)

    return 
}