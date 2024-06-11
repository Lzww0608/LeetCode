const MOD int = 1e9 + 7
func numberOfArrays(s string, k int) int {
    n := len(s)
    f := make([]int, n + 1)
    f[0] = 1
    
    for i := range s {
        t, add := 0, 1
        for j := i; j >= 0 && i - j <= 10; j-- {
            t += int(s[j] - '0') * add
            if t > k {
                break
            }
            if s[j] != '0' { 
                f[i+1] = (f[i+1] + f[j]) % MOD
            }
            add *= 10
        }
    }

    return f[n]
}
