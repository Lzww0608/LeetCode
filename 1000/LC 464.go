func canIWin(a int, b int) bool {
    if a * (a + 1) / 2 < b {
        return false
    }
    if b <= 1 {
        return true
    }
    memo := make(map[int]bool)
    var dfs func(int, int) bool
    dfs = func(mask, now int) bool {
        if now >= b {
            return false
        }
        if v, ok := memo[mask]; ok {
            return v
        }
        for i := 1; i <= a; i++ {
            cur := 1 << i
            if mask & cur == 0 {
                if !dfs(mask | cur, now + i) {
                    memo[mask] = true
                    return true
                }
            }
        }
        memo[mask] = false
        return false
    }
    return dfs(0, 0)
}
