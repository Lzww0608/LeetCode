func distinctEchoSubstrings(s string) (ans int) {
    n := len(s)
    m := map[string]bool{}

    for i := range s {
        for j := i + 1; j < n; j++ {
            l := j - i 
            if j * 2 - i <= n && s[i:i+l] == s[j:j+l] && !m[s[i:i+l]] {
                m[s[i:i+l]] = true
                ans++
            }
        }
    }

    return 
}
