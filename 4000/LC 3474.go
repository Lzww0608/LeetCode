func generateString(s string, t string) string {
    n, m := len(s), len(t)
    ans := make([]byte, n + m - 1)
    for i := range ans {
        ans[i] = '?'
    }

    for i := range s {
        if s[i] != 'T' {
            continue
        }
        for j := range t {
            if ans[i + j] != '?' && ans[i + j] != t[j] {
                return ""
            } 
            ans[i + j] = t[j]
        }
    }

    for i := range s {
        if s[i] != 'F' {
            continue
        }
        if slices.Equal(ans[i : i + m], []byte(t)) {
            return ""
        }
    }

    old := slices.Clone(ans)
    for i := range ans {
        if ans[i] == '?' {
            ans[i] = 'a'
        }
    }

    for i := range s {
        if s[i] != 'F' {
            continue
        }
        f := true
        if slices.Equal(ans[i : i + m], []byte(t)) {
            f = false
            for j := m - 1; j >= 0; j-- {
                if old[i + j] == '?' {
                    ans[i + j] = 'b'
                    f = true
                    break
                }
            }
        }
        if !f {
            return ""
        }
    }

    return string(ans)
}