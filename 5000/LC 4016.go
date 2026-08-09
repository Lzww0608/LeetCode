func maxArea(mat [][]int) (ans int) {
    m, n := len(mat), len(mat[0])
    f := make([][]int, m + 1)
    ff := make([][]int, m + 1)
    g := make([][]int, m + 1)
    gg := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
        g[i] = make([]int, n + 1)
        ff[i] = make([]int, n + 1)
        gg[i] = make([]int, n + 1)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 {
                f[i+1][j+1] =
                    min(f[i][j], f[i+1][j], f[i][j+1]) + 1
            }

            ff[i+1][j+1] =
                max(f[i+1][j+1], ff[i+1][j], ff[i][j+1])
        }
    }

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            if mat[i][j] == 1 {
                g[i][j] =
                    min(g[i+1][j+1], g[i+1][j], g[i][j+1]) + 1
            }

            gg[i][j] =
                max(g[i][j], gg[i+1][j], gg[i][j+1])
        }
    }

    for i := 1; i < m; i++ {
        k := min(ff[i][n], gg[i][0])
        ans = max(ans, k*k)
    }

    for j := 1; j < n; j++ {
        k := min(ff[m][j], gg[0][j])
        ans = max(ans, k*k)
    }

    return
}