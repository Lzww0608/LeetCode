func numOfWays(n int) (ans int) {
    const MOD int = 1e9 + 7

    f := [3][3][3]int{}
    g := [3][3][3]int{}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            for k := 0; k < 3; k++ {
                if i != j && j != k {
                    g[i][j][k] = 1
                }
            }
        }
    }

    for t := 2; t <= n; t++ {
        for i := 0; i < 3; i++ {
            for j := 0; j < 3; j++ {
                for k := 0; k < 3; k++ {
                    for p := 0; p < 3; p++ {
                        for q := 0; q < 3; q++ {
                            for r := 0; r < 3; r++ {
                                if i != j && j != k && i != p && j != q && k != r {
                                    f[i][j][k] += g[p][q][r]
                                    f[i][j][k] %= MOD
                                }
                            }
                        }
                    }
                }
            }
        }
        g, f = f, [3][3][3]int{}
    }

    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            for k := 0; k < 3; k++ {
                if i != j && j != k {
                    ans = (ans + g[i][j][k]) % MOD
                }
            }
        }
    }
    
    return 
}


func numOfWays(n int) int {
    const MOD int = 1e9 + 7

    a, b := 6, 6
    for i := 2; i <= n; i++ {
        a, b = a * 3 + b * 2, a * 2 + b * 2
        a %= MOD
        b %= MOD
    }

    return (a + b) % MOD
}