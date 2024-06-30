func isValidSerialization(preorder string) bool {
    n, idx := len(preorder), 0

    var dfs func() bool
    dfs = func() bool {
        if idx >= n {
            return false
        }
        if preorder[idx] != '#' {
            for idx < n && preorder[idx] >= '0' && preorder[idx] <= '9' {
                idx++
            }
            idx++
            return dfs() && dfs()
        } 
        idx += 2
        return true
    }

    return dfs() && idx >= n
}
