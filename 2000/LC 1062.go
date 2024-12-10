const MOD int = 1_070_777_777
func longestRepeatingSubstring(s string) (ans int) {
    n := len(s)
    base := 9e8 - rand.Intn(1e8)
    powBase := make([]int, n + 1)
    preHash := make([]int, n + 1)
    powBase[0] = 1
    for i, c := range s {
        powBase[i+1] = (powBase[i] * base) % MOD
        preHash[i+1] = (preHash[i] * base + int(c)) % MOD
    } 

    // [l, r)
    subHash := func(l, r int) int {
        return ((preHash[r] - preHash[l] * powBase[r-l]) % MOD + MOD) % MOD
    }

    check := func(d int) bool {
        m := map[int]bool{}
        for i := 0; i + d <= n; i++ {
            hash := subHash(i, i + d)
            if m[hash] {
                return true
            }
            m[hash] = true
        }
        return false
    }

    l, r := 1, n
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            ans = mid
            l = mid + 1
        } else {
            r = mid
        }
    }

    return 
}