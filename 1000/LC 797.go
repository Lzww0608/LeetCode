func allPathsSourceTarget(graph [][]int) (ans [][]int) {
    n := len(graph)

    var dfs func(int, int, []int)
    dfs = func(u, mask int, cur []int) {
        if u == n - 1 {
            ans = append(ans, append([]int(nil), cur...))
            return
        }
        
        for _, v := range graph[u] {
            if mask & (1 << v) == 0 {
                cur = append(cur, v)
                dfs(v, mask | (1 << v), cur)
                cur = cur[:len(cur)-1]
            }
        }
    }
    dfs(0, 0, []int{0})

    return 
}