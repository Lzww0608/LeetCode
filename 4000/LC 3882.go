const N int = 1 << 10
func minCost(a [][]int) int {
    m, n := len(a), len(a[0])
    f := make([][][N]bool, m + 1)
    for i := range f {
        f[i] = make([][N]bool, n + 1)
    }
    f[0][1][0] = true

    for i := range m {
        for j, x := range a[i] {
            for k := range N {
                t := k ^ x 
                if f[i][j + 1][k] || f[i + 1][j][k] {
                    f[i + 1][j + 1][t] = true
                }
            }
        }
    }

    for i := range N {
        if f[m][n][i] {
            return i
        }
    }

    return -1
}