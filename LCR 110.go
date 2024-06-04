func allPathsSourceTarget(graph [][]int) (ans [][]int) {
    path := []int{}
    n := len(graph)

    var dfs func(int) 
    dfs = func(idx int) {
        path = append(path, idx)
        if idx == n - 1 {
            ans = append(ans, append([]int(nil), path...))
            path = path[:len(path)-1]
            return
        }

        for _, x := range graph[idx] {
            dfs(x)
        }
        path = path[:len(path)-1]
    }
    dfs(0)

    return
}