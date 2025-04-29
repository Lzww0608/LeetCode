const MOD int = 1_070_777_777
func maxProduct(s string) int64 {
    n := len(s)
    base := 9e8 - rand.Intn(1e8)
    powBase := make([]int, n + 1)
    preHash := make([]int, n + 1)
    sufHash := make([]int, n + 1)
    powBase[0] = 1

    for i, j := 1, n; i <= n; i, j = i + 1, j - 1 {
        preHash[i] = (preHash[i-1] * base + int(s[i-1])) % MOD
        sufHash[i] = (sufHash[i-1] * base + int(s[j-1])) % MOD  
        powBase[i] = powBase[i-1] * base % MOD
    }

    getPre := func(l, r int) int {
        return ((preHash[r] - preHash[l-1] * powBase[r-l+1]) % MOD + MOD) % MOD
    }

    getSuf := func(l, r int) int {
        return ((sufHash[r] - sufHash[l-1] * powBase[r-l+1]) % MOD + MOD) % MOD
    }

    w := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        l, r := 0, min(i - 1, n - i) 
        for l < r {
            mid := (l + r + 1) >> 1
            if getPre(i - mid, i) == getSuf(n + 1 - i - mid, n + 1 - i) {
                l = mid
            } else {
                r = mid - 1
            }
        }
        w[i] = l
    }

    f := make([]int, n + 1)
    g := make([]int, n + 1)
    for i, j, mx := 1, 1, 0; i <= n; i++ {
        for j <= n && j + w[j] <= i {
            mx = max(mx, w[j])
            j++
        }
        mx = max(mx, i - j)
        f[i] = mx 
    }

    for i, j, mx := n, n, 0; i > 0; i-- {
        for j - w[j] >= i {
            mx = max(mx, w[j])
            j--
        }
        mx = max(mx, j - i)
        g[i] = mx
    }

    ans := 0
    for i := 1; i < n; i++ {
        ans = max(ans, (f[i] * 2 + 1) * (g[i+1] * 2 + 1))
    }

    return int64(ans)
}