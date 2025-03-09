func minimumMoves(arr []int) int {
    n := len(arr)
    f := make([][]int, n)
    g := make([][]int, n)
    for i := range f {
        f[i] = make([]int, n)
        f[i][i] = 1
        x := arr[i]
        for j := i; j < n; j++ {
            if arr[j] == x {
                g[i] = append(g[i], j)
            }
        }
    }

    for i := n - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            if i == j - 1 {
                if arr[i] == arr[j] {
                    f[i][j] = 1
                } else {
                    f[i][j] = 2
                }
            } else {
                if arr[i] == arr[j] {
                    f[i][j] = f[i+1][j-1]
                } else {
                    f[i][j] = f[i+1][j-1] + 2
                }
                for _, k := range g[i] {
                    if k >= j {
                        break
                    } else {
                        f[i][j] = min(f[i][j], f[i][k] + f[k+1][j])
                    }
                }
            }
        }
    }

    return f[0][n-1]
}