func maximumXor(s string, t string) string {
    n := len(t)
    a, b := 0, 0
    for _, c := range t {
        if c == '0' {
            a++
        } else {
            b++
        }
    }

    ans := make([]byte, n)
    for i := range s {
        if s[i] == '0' {
            if b > 0 {
                ans[i] = '1'
                b--
            } else {
                ans[i] = '0'
                a--
            }
        } else {
            if a > 0 {
                ans[i] = '1'
                a--
            } else {
                ans[i] = '0'
                b--
            }
        }
    }

    return string(ans)
}