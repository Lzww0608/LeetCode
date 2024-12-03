const MOD int = 1_000_000_007
func distinctSequences(n int) (ans int) {
    if n == 1 {
        return 6
    }
    f := [6][6]int{}
    for i := 1; i <= 6; i++ {
        for j := 1; j <= 6; j++ {
            if i != j && gcd(i, j) == 1 {
                f[i-1][j-1]++
            }
        }const MOD int = 1_000_000_007
func distinctSequences(n int) (ans int) {
    if n == 1 {
        return 6
    }
    f := [6][6]int{}
    for i := 1; i <= 6; i++ {
        for j := 1; j <= 6; j++ {
            if i != j && gcd(i, j) == 1 {
                f[i-1][j-1]++
            }
        }
    }

    for t := 3; t <= n; t++ {
        g := [6][6]int{}
        for i := 1; i <= 6; i++ {
            for j := 1; j <= 6; j++ {
                if i == j || gcd(i, j) != 1 {
                    continue
                }
                for k := 1; k <= 6; k++ {
                    if k != i && k != j && gcd(k, j) == 1 {
                        g[k-1][j-1] = (g[k-1][j-1] + f[j-1][i-1]) % MOD
                    }
                }
            }
        }
        f = g
    }

    for _, v := range f {
        for _, x := range v {
            ans = (ans + x) % MOD
        }
    }

    return
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y 
    }
    return x
}
    }

    for t := 3; t <= n; t++ {
        g := [6][6]int{}
        for i := 1; i <= 6; i++ {
            for j := 1; j <= 6; j++ {
                if i == j || gcd(i, j) != 1 {
                    continue
                }
                for k := 1; k <= 6; k++ {
                    if k != i && k != j && gcd(k, j) == 1 {
                        g[k-1][j-1] = (g[k-1][j-1] + f[j-1][i-1]) % MOD
                    }
                }
            }
        }
        f = g
    }

    for _, v := range f {
        for _, x := range v {
            ans = (ans + x) % MOD
        }
    }

    return
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y 
    }
    return x
}