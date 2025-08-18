const MOD int = 1_000_000_007
func countKReducibleNumbers(s string, k int) (ans int) {
    if s == "1" {
        return 0
    }

    n := len(s)
    f := make([]int, n + 1)
    for i := 2; i <= n; i++ {
        j := bits.OnesCount(uint(i))
        f[i] = f[j] + 1
    }

    C := make([][]int, n + 1)
    for i := 0; i <= n; i++ {
        C[i] = make([]int, n + 1)
        C[i][0] = 1
        for j := 1; j <= i; j++ {
            C[i][j] = (C[i - 1][j - 1] + C[i - 1][j]) % MOD
        }
    } 

    for i := 1; i <= n; i++ {
        if f[i] > k - 1 {
            continue
        }

        cnt := i
        for j, c := range s {
            if c == '1' {
                ans = (ans + C[n - j - 1][cnt]) % MOD
                cnt--
                if cnt < 0 {
                    break
                }
            }
        }
    }

    return
}