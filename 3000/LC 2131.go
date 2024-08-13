func longestPalindrome(words []string) (ans int) {
    m := map[string]int{}
    pal := 0
    for _, s := range words {
        if m[s] > 0 {
            ans += len(s) * 2
            m[s]--
            continue
        }
        t := []byte(s)
        n := len(s)
        for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
            t[i], t[j] = t[j], t[i]
        }
        rev := string(t)
        m[rev]++
    }

    for s := range m {
        if m[s] == 0 {
            continue
        }
        n := len(s)
        f := true
        for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1 {
            if s[i] != s[j] {
                f = false
                break
            }
        }
        if f {
            pal = max(pal, n)
        }
    }

    return ans + pal
}