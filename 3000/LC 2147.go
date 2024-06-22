const MOD int = 1e9 + 7
func numberOfWays(s string) int {
    cnt, ans := 0, 1
    n := len(s)
    for i := 0; i < n; i++ {
        if s[i] == 'S' {
            cnt++
        }
        if cnt > 0 && cnt & 1 == 0 {
            i++
            k := 0
            for i < n && s[i] != 'S' {
                k++
                i++
            }
            if i < n {
                ans = (ans * (k + 1)) % MOD
            } 
            i--
        }
    }
    if cnt == 0 || cnt & 1 == 1 {
        return 0
    }
    return ans
}
