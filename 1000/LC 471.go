func encode(s string) string {
    n := len(s)
    f := make([][]string, n)
    for i := range f {
        f[i] = make([]string, n)
        f[i][i] = s[i:i+1]
    }

    for i := n - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            f[i][j] = s[i:j+1]
            ss := s[i:j+1]
            if len(ss) < 5 {
                continue
            }
            t := ss + ss
            idx := strings.Index(t[1:], ss) + 1
            cnt := len(ss) / idx
            tmp := strconv.Itoa(cnt) + "[" + f[i][i+idx-1] + "]" 
            if len(f[i][j]) >= len(tmp) {
                f[i][j] = tmp
            }
            for k := i + 1; k < j; k++ {
                if len(f[i][j]) > len(f[i][k]) + len(f[k+1][j]) {
                    f[i][j] = f[i][k] + f[k+1][j]
                }
            }

        }
    }

    return f[0][n-1]
}