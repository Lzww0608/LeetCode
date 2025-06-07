func possiblyEquals(s string, t string) bool {
    n, m := len(s), len(t)
    f := make([][]map[int]bool, n + 1)
    for i := range f {
        f[i] = make([]map[int]bool, m + 1)
        for j := range f[i] {
            f[i][j] = make(map[int]bool)
        }
    }
    f[0][0][0] = true;

    for i := 0; i <= n; i++ {
        for j := 0; j <= m; j++ {
            for delta := range f[i][j] {
                num := 0
                if delta <= 0 {
                    for p := i; p < min(i + 3, n); p++ {
                        if s[p] >= '0' && s[p] <= '9' {
                            num = num * 10 + int(s[p] - '0')
                            f[p + 1][j][delta + num] = true;
                        } else {
                            break
                        }
                    }
                }

                num = 0
                if delta >= 0 {
                    for p := j; p < min(j + 3, m); p++ {
                        if t[p] >= '0' && t[p] <= '9' {
                            num = num * 10 + int(t[p] - '0')
                            f[i][p + 1][delta - num] = true;
                        } else {
                            break
                        }
                    }
                }

                if i < n && delta < 0 && (s[i] < '0' || s[i] > '9') {
                    f[i + 1][j][delta + 1] = true;
                }

                if j < m && delta > 0 && (t[j] < '0' || t[j] > '9') {
                    f[i][j + 1][delta - 1] = true;
                }

                if i < n && j < m && delta == 0 && s[i] == t[j] {
                    f[i+1][j+1][0] = true;
                }
            }
        }
    }

    return f[n][m][0]
}