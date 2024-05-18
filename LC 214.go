func shortestPalindrome(s string) string {
    n := len(s)
    if n == 0 || n == 1 {
        return s
    }
    check := func(i, j int) bool {
        for i < j {
            if s[i] != s[j] {
                return false
            }
            i++
            j--
        }
        return true
    }

    builder := strings.Builder{}
    for i := n - 1; i >= 0; i-- {
        if check(0, i) {
            return builder.String() + s
        }
        builder.WriteByte(s[i])
    }

    return s
}