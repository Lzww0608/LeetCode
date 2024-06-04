const MOD int  = 1e9 + 7
func numWays(steps int, arrLen int) int {
    if arrLen == 1 {
        return 1
    }
    n := min(steps / 2 + 1, arrLen)
    f := make([]int, n)
    g := make([]int, n)
    g[0] = 1

    for i := 1; i <= steps; i++ {  
        for j := 0; j < n; j++ {
            f[j] = g[j] % MOD 
            if j - 1 >= 0 {
                f[j] = (f[j] + g[j-1]) % MOD  
            }
            if j + 1 < n {
                f[j] = (f[j] + g[j+1]) % MOD  
            }
        }
        g, f = f, g  
    }

    return g[0]
}
