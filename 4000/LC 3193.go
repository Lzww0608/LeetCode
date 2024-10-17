const MOD int = 1_000_000_007

func numberOfPermutations(n int, requirements [][]int) (ans int) {
    req := make([]int, n)
    for i := range req {
        req[i] = -1
    }

    for _, v := range requirements {
        req[v[0]] = v[1]
    }

    m := req[n-1] + 1
    f := make([][]int, n)
    for i := range f {
        f[i] = make([]int, m)
    }
    f[0][0] = 1
    
    for i := 0; i < n - 1; i++ {
        if req[i] != -1 {
            for j := 0; j < m; j++ {
                if req[i] != j {
                    f[i][j] = 0
                }
            }
        }
        
        for k := 0; k <= i + 1; k++ {
            for j := 0; j + k < m; j++ {
                f[i+1][j+k] += f[i][j]
                f[i+1][j+k] %= MOD
            }
        }
    }

    return f[n-1][m-1]
}