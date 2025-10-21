func processStr(s string) string {
    ans := []byte{}
    for i := range s {
        if s[i] == '*' {
            if len(ans)> 0 {
                ans = ans[:len(ans) - 1]
            }
        } else if s[i] == '#' {
            ans = append(ans, ans...)
        } else if s[i] == '%' {
            slices.Reverse(ans)
        } else {
            ans = append(ans, s[i])
        }
    }

    return string(ans)
}