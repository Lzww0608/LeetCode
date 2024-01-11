func wordBreak(s string, wordDict []string) []string {
    m := map[string]bool{}
    n := len(s)
    for _, x := range wordDict {
        m[x] = true
    }
    ans, res := []string{}, ""
    var dfs func(int)
    dfs = func(idx int) {
        if idx >= n {
            ans = append(ans, res[:len(res)-1])
            return
        }
        var t string
        for i := idx; i < n; i++ {
            t += string(s[i])
            if m[t] {
                res += t + " "
                dfs(i+1)
                res = res[:len(res)-1-len(t)]
            }
        }
    }
    dfs(0)
    return ans
}