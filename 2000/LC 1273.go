func deleteTreeNodes(nodes int, parent []int, value []int) int {
    g := make([][]int, nodes)
    for i := 1; i < nodes; i++ {
        g[parent[i]] = append(g[parent[i]], i)
    }

    var dfs func(int) (int, int) 
    dfs = func(u int) (sum, cnt int) {
        for _, v := range g[u] {
            a, b := dfs(v)
            sum += a 
            cnt += b 
        }
        sum += value[u]
        cnt += 1 
        if sum == 0 {
            cnt = 0
        }
        return 
    }

    _, ans := dfs(0)
    return ans
}