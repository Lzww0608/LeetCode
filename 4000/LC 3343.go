const MOD int = 1_000_000_007

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        r >>= 1 
        a = a * a % MOD
    }

    return res
}

func fac(x int) int {
    if x == 0 {
        return 1
    }

    return fac(x - 1) * x % MOD
}

func countBalancedPermutations(s string) int {
    n := len(s)
    cnt := [10]int{}
    sum := 0
    for i := range s {
        x := int(s[i] - '0')
        cnt[x]++
        sum += x
    }
    if sum & 1 == 1 {
        return 0
    }

    f := make([][]int, n / 2 + 1)
    for i := range f {
        f[i] = make([]int, sum / 2 + 1)
    }
    f[0][0] = 1 
    for i := range s {
        x := int(s[i] - '0')
        g := make([][]int, n / 2 + 1)
        for j := range g {
            g[j] = make([]int, sum / 2 + 1)
        }

        for j := 0; j <= n / 2; j++ {
            for k := sum / 2; k >= 0; k-- {
                g[j][k] = f[j][k]
                if j > 0 && k >= x {
                    g[j][k] = (g[j][k] + f[j-1][k-x]) % MOD 
                }
            }
        }
        f = g
    }
    
    ans := fac(n / 2) * fac(n - n / 2) % MOD * f[n / 2][sum / 2] % MOD
    for _, x := range cnt {
        ans = ans * quickPow(fac(x), MOD - 2) % MOD
    }

    return ans
}