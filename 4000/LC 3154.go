func waysToReachStair(k int) int {
    type pair struct {
        i, j int
        k bool
    }
    m := map[pair]int{}

    var dfs func(int, int, bool) int 
    dfs = func(i, j int, down bool) (res int) {
        if i < 0 || i > k + 1 {
            return 0
        }

        if v, ok := m[pair{i, j, down}]; ok {
            return v
        }
        res += dfs(i + (1 << j), j + 1, false)
        if !down {
            res += dfs(i - 1, j, true)
        }
        if i == k {
            res++
        }
        m[pair{i, j, down}] = res
        return
    }

    return dfs(1, 0, false)
}