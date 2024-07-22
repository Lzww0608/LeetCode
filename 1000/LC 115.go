const MOD int = 1_000_000_007
func numDistinct(s string, t string) int {
    n, m := len(s), len(t)
    f := make([]int, m + 1)

    for i := 0; i < n; i++ {
        pre := 1
        for j := 0; j < m; j++ {
            cur := f[j+1]
            if s[i] == t[j] {
                f[j+1] += pre
            } 
            pre = cur
        }
    }

    return f[m]
}