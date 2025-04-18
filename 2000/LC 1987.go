const MOD int = 1_000_000_007
func numberOfUniqueGoodSubsequences(s string) int {
    n := len(s)
    f := [2]int{}
    for i := n - 1; i >= 0; i-- {
        x := int(s[i] - '0')
        f[x] = (f[0] + f[1] + 1) % MOD
    }

    ans := f[1]
    if strings.Contains(s, "0") {
        ans = (ans + 1) % MOD
    }

    return ans
}