func canCross(stones []int) bool {
    n := len(stones)
    m := make([]map[int]bool, n)
    for i := range m {
        m[i] = make(map[int]bool)
    }

    var dfs func(int, int) bool
    dfs = func(i, last int) bool {
        if i == n - 1 {
            return true
        }
        if v, ok := m[i][last]; ok {
            return v
        }
        for j := last  - 1;j <= last + 1; j++ {
            if j <= 0 {
                continue
            }
            pos := sort.SearchInts(stones, stones[i] + j)
            if pos != n && stones[pos] == stones[i] + j && dfs(pos, j) {
                m[i][last] = true
                return true
            }
        }
        m[i][last] = false
        return false
    }
    return dfs(0, 0)
}