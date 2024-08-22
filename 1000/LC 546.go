func removeBoxes(boxes []int) int {
    n := len(boxes)
    f := make([][][]int, n)
    for i := range f {
        f[i] = make([][]int, n)
        for j := range f[i] {
            f[i][j] = make([]int, n)
        }
    }

    var dfs func(int, int, int) int
    dfs = func(i, j, k int) int {
        if i > j {
            return 0
        }
        if f[i][j][k] != 0 {
            return f[i][j][k]
        }

        f[i][j][k] = dfs(i, j - 1, 0) + (k + 1) * (k + 1)
        for p := i; p < j; p++ {
            if boxes[p] == boxes[j] {
                f[i][j][k] = max(f[i][j][k], dfs(i, p, k + 1) + dfs(p + 1, j - 1, 0))
            }
        }

        return f[i][j][k]
    }

    return dfs(0, n - 1, 0)
}