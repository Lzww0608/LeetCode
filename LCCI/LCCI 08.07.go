func permutation(str string) (ans []string) {
    s := []byte(str)
    n := len(s)

    var dfs func(int)
    dfs = func(idx int) {
        if idx == n {
            ans = append(ans, string(s))
            return
        }

        for i := idx; i < n; i++ {
            s[i], s[idx] = s[idx], s[i]
            dfs(idx + 1)
            s[i], s[idx] = s[idx], s[i]
        }
    }
    dfs(0)

    return
}