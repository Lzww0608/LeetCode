func maxPathScore(grid [][]int, k int) int {
    m, n := len(grid), len(grid[0])
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, k + 1)
        for j := range k + 1 {
            f[i][j] = -1
        }
    }
    f[0][k] = 0

    for i := range m {
        for j := range n {
            x := grid[i][j]
            c := 1
            if x == 0 {
                c = 0
            }

            tmp := make([]int, k + 1)
            for t := range tmp {
                tmp[t] = -1 
            }
            for t := k; t >= c; t-- {
                if f[j][t] != -1 {
                    tmp[t - c] = max(tmp[t - c], f[j][t] + x)
                }
                if j > 0 && f[j - 1][t] != -1 {
                    tmp[t - c] = max(tmp[t - c], f[j - 1][t] + x)
                }
            }
            f[j] = tmp
        }
    }

    return slices.Max(f[n - 1])
}