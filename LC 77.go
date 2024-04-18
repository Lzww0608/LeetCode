func combine(n int, k int) (ans [][]int) {
    path := []int{}

    var dfs func(int)
    dfs = func(idx int) {
        if len(path) == k {
            ans = append(ans, append([]int(nil), path...))
            return 
        }
        for i := idx; i <= n - (k - len(path)) + 1; i++ {
            path = append(path, i)
            dfs(i + 1)
            path = path[:len(path)-1]
        }
    }
    dfs(1)

    return 
}