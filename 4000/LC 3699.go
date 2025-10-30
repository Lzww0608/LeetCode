const MOD int = 1_000_000_007
func zigZagArrays(n int, l int, r int) int {
    k := r - l + 1
    f := make([]int, k)
    for i := range f {
        f[i] = 1
    }

    for i := range n - 1 {
        if i & 1 == 0 {
            pre := 0
            for j, x := range f {
                f[j] = pre % MOD
                pre += x
            }
        } else {
            suf := 0 
            for j := k - 1; j >= 0; j-- {
                x := f[j]
                f[j] = suf % MOD
                suf += x
            }
        }
    }

    ans := 0
    for _, x := range f {
        ans = (ans + x) % MOD
    }

    return ans * 2 % MOD
}