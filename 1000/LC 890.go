func findAndReplacePattern(words []string, pattern string) []string {
    ans := []string{}
    for _, s := range words {
        m1 := map[byte]byte{}
        m2 := map[byte]byte{}
        f := true
        for i := range s {
            a, ok1 := m1[s[i]]
            b, ok2 := m2[pattern[i]]
            if !ok1 && !ok2 {
                m1[s[i]] = pattern[i]
                m2[pattern[i]] = s[i]
            } else if a != pattern[i] || b != s[i] {
                f = false
                break
            }
        }
        if f {
            ans = append(ans, s)
        }
    }
    return ans
}
