func splitString(s string) bool {
    n := len(s)

    var dfs func(int, int) bool
    dfs = func(idx, pre int) bool {
        if idx >= n {
            return true
        }
        cur := 0
        for i := idx; i < n; i++ {
            cur = cur * 10 + int(s[i] - '0')
            if cur >= pre {
                break
            }
            if cur == pre - 1 && dfs(i + 1, cur) {
                return true
            }
        }

        return false
    }

    cur := 0
    for i := 0; i < n - 1; i++ {
        cur = cur * 10 + int(s[i] - '0')
        if dfs(i + 1, cur) {
            return true
        }
    }
    return false
}