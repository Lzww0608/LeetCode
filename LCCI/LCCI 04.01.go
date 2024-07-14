func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
    g := make([][]int, n)
    for _, e := range graph {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
    }
    vis := make([]bool, n)

    var dfs func(int) bool
    dfs = func(cur int) bool {
        if vis[cur] {
            return false
        } else if vis[cur] = true; cur == target {
            return true
        }

        for _, x := range g[cur] {
            if dfs(x) {
                return true
            }
        }
        
        return false
    }

    return dfs(start)
}