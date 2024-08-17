func braceExpansionII(expression string) (ans []string) {
    m := map[string]bool{}

    var dfs func(string)
    dfs = func(s string) {
        j := strings.Index(s, "}")
        if j == -1 {
            m[s] = true
            return
        }
        i := strings.LastIndex(s[:j], "{")
        a, c := s[:i], s[j+1:]
        for _, b := range strings.Split(s[i+1:j], ",") {
            dfs(a + b + c)
        }
    }
    dfs(expression)

    for s := range m {
        ans = append(ans, s)
    }
    sort.Strings(ans)

    return
}