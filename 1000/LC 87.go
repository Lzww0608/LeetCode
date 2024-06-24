func isScramble(s1 string, s2 string) bool {
    a := []byte(s1)
    b := []byte(s2)
    n := len(a)

    f := make([][][]bool, n)
    for i := range f {
        f[i] = make([][]bool, n)
        for j := range f[i] {
            f[i][j] = make([]bool, n + 1)
            if a[i] == b[j] {
                f[i][j][1] = true
            }
        }
    }

    for d := 2; d <= n; d++ {
        for i := 0; i + d <= n; i++ {
            for j := 0; j + d <= n; j++ {
                for k := 1; k < d; k++ {
                    if f[i][j][k] && f[i+k][j+k][d-k] {
                        f[i][j][d] = true
                        break
                    }

                    if f[i][j+d-k][k] && f[i+k][j][d-k] {
                        f[i][j][d] = true
                        break
                    }
                }
            }
        }
    }

    return f[0][0][n]
}