func longestDiverseString(a int, b int, c int) string {
    ans := make([]byte, 0, a + b + c)
    for a > 0 || b > 0 || c > 0 {
        mx := max(a, b, c)
        if a == mx {
            if len(ans) >= 2 && ans[len(ans)-1] == 'a' && ans[len(ans)-2] == 'a' {
                if b >= c && b > 0 {
                    b--
                    ans = append(ans, 'b')
                } else if c > 0 {
                    c--
                    ans = append(ans, 'c')
                } else {
                    break
                }
            } else {
                ans = append(ans, 'a')
                a--
            }
        } else if b == mx {
            if len(ans) >= 2 && ans[len(ans)-1] == 'b' && ans[len(ans)-2] == 'b' {
                if a >= c && a > 0 {
                    a--
                    ans = append(ans, 'a')
                } else if c > 0 {
                    c--
                    ans = append(ans, 'c')
                } else {
                    break
                }
            } else {
                ans = append(ans, 'b')
                b--
            }
        } else {
            if len(ans) >= 2 && ans[len(ans)-1] == 'c' && ans[len(ans)-2] == 'c' {
                if a >= b && a > 0 {
                    a--
                    ans = append(ans, 'a')
                } else if b > 0 {
                    b--
                    ans = append(ans, 'b')
                } else {
                    break
                }
            } else {
                ans = append(ans, 'c')
                c--
            }
        }
    }

    return string(ans)
}