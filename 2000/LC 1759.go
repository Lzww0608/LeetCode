const MOD int = 1e9 + 7
func countHomogenous(s string) (ans int) {
    n := len(s)
    l := 0
    for l < n {
        r := l + 1
        ans++
        for r < n && s[r] == s[l] {
            r++
            ans = (ans + r - l) % MOD
        }
        l = r
    }

    return 
}
