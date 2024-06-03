var MOD int = int(1e9 + 7)
type matrix [5][5]int

func mul(a, b matrix) matrix {
    c := matrix{}
    
    for i := range a {
        for j := range a[i] {
            for k, v := range a[i] {
                c[i][j] = (c[i][j] + v * b[k][j]) % MOD
            }
        }
    }

    return c
}

func countVowelPermutation(n int) int {
    ans := matrix {
        {1, 1, 1, 1, 1},
    }

    mat := matrix {
        {0, 1, 0, 0, 0},
        {1, 0, 1, 0, 0},
        {1, 1, 0, 1, 1},
        {0, 0, 1, 0, 1},
        {1, 0, 0, 0, 0},
    }

    for k := n - 1; k > 0; k >>= 1 {
        if k & 1 == 1 {
            ans = mul(ans, mat)
        }
        mat = mul(mat, mat)
    }

    res := 0
    for i := range ans {
        res += ans[0][i]
    }

    return res % MOD
}
