const MOD int64 = 1e9 + 7
func numSub(s string) (ans int) {
    n := len(s)
    l, r := -1, 0
    for r <= n {
        if r == n || s[r] != '1' {
            ans += int(int64(r - l) * int64(r - l - 1) / 2 % MOD) 
            //fmt.Println(l, r)
            l = r
        }
        r++
    }

    return 
}