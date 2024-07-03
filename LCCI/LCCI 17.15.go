func longestWord(words []string) string {
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) > len(words[j]) || (len(words[i]) == len(words[j]) && words[i] < words[j])
    })

    m := map[string]struct{}{}
    for _, s := range words {
        m[s] = struct{}{}
    }

    var dfs func(string) bool
    dfs = func(s string) bool {
        if s == "" {
            return true
        }
        for i := range s {
            if _, ok := m[s[:i+1]]; ok {
                if dfs(s[i+1:]) {
                    return true
                }
            }
        }
        return false
    }

    for _, s := range words {
        delete(m, s)
        if dfs(s) {
            return s
        }
    }

    return ""
}