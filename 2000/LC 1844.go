func replaceDigits(s string) string {
    ans := []byte(s)
    for i := 1; i < len(s); i += 2 {
        c := byte(s[i] - '0' + s[i-1])
        if c > 'z' {
            c = byte('a' + (c - 'z'))
        }
        ans[i] = c
    }

    return string(ans)
}