type matrix [2][2]int

func mul(a, b matrix) matrix {
    c := matrix{}
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            for k, v := range a[i] {
                c[i][j] += v * b[k][j]
            }
        }
    }

    return c
}

func fib(n int) int {
    if n < 2 {
        return n
    }

    
    mat := matrix {
        {1, 1},
        {1, 0},
    }

    ans := matrix {
        {1, 0},
        {0, 1},
    }

    for k := n - 1; k > 0; k >>= 1 {
        if k & 1 == 1 {
            ans = mul(ans, mat)
        }
        mat = mul(mat, mat)
    }
    return ans[0][0]
}

