func removeStars(s string) string {
    ans := make([]byte, 0, len(s))
    for i := range s {
        if s[i] != '*' {
            ans = append(ans, s[i])
        } else {
            ans = ans[:len(ans)-1]
        }
    }

    return string(ans)
}
