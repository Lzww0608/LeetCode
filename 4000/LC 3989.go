func maxConsistentColumns(a [][]int, limit int) int {
    m, n := len(a), len(a[0])
    f := make([]int, n)
    f[0] = 1

    for j := 1; j < n; j++ {
        f[j] = 1
        for k := j - 1; k >= 0; k-- {
            b := true 
            for i := range m {
                if abs(a[i][k] - a[i][j]) > limit {
                    b = false
                    break
                } 
            }

            if b {
                f[j] = max(f[j], f[k] + 1)
            }
        }
    }

    return slices.Max(f)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}