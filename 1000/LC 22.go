func generateParenthesis(n int) (ans []string) {
    var s string
    
    var dfs func(int, int)
    dfs = func(a, b int) {
        if a == 0 && b == 0 {
            ans = append(ans, s)
            return
        }
        if a == b {
            s += "("
            dfs(a - 1, b)
            s = s[:len(s)-1]
        } else {
            if a > 0 {
                s += "("
                dfs(a - 1, b)
                s = s[:len(s)-1]
            }
            s += ")"
            dfs(a, b - 1)
            s = s[:len(s)-1]
        }
    }
    dfs(n, n)

    return 
}
