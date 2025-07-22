func tourOfKnight(m int, n int, r int, c int) [][]int {
    dirs := [][]int{{1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}}

    f := make([][]int, m)
    for i := range f {
        f[i] = make([]int, n)
        for j := range f[i] {
            f[i][j] = -1
        }
    }

    find := false

    var dfs func(int, int, int)
    dfs = func(i, j, cur int) {
        if i < 0 || i >= m || j < 0 || j >= n || f[i][j] != -1 {
            return
        }
        f[i][j] = cur
        if cur == m * n - 1 {
            find = true
        }
        if find {
            return
        }

        for _, dir := range dirs {
            x, y := i + dir[0], j + dir[1]
            dfs(x, y, cur + 1)
        }
        if !find {
            f[i][j] = -1
        }
        return
    }
    dfs(r, c, 0)

    return f
}