func makeGood(s string) string {
    n := len(s)
    ans := make([]byte, 0, n)
    for i := range s {
        if len(ans) == 0 || ans[len(ans)-1] + 32 != s[i] && ans[len(ans)-1] - 32 != s[i] {
            ans = append(ans, s[i])
        } else {
            ans = ans[:len(ans)-1]
        }
    }

    return string(ans)
}