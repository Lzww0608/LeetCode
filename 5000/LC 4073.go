const MOD int = 1_000_000_007
type Matrix [][]int
func countGoodStrings(n int64) int {
    m := Matrix {
        {1, 1},
        {1, 0},
    }

    f := Matrix {
        {1},
        {0},
    }

    f = m.pow(n - 1, f)
    return f[0][0] * 2 % MOD
}

func (m Matrix) pow(r int64, f Matrix) Matrix {
    res := f 
    for r > 0 {
        if r & 1 == 1 {
            res = m.mul(res)
        }

        m = m.mul(m)
        r >>= 1
    }

    return res
}


func (a Matrix) mul(b Matrix) Matrix {
    n, m := len(a), len(b)
    c := make(Matrix, n)
    for i := range c {
        c[i] = make([]int, m)
    }

    for i := range n {
        for k := range a[i] {
            if a[i][k] == 0 {
                continue
            }

            for j := range b[k] {
                c[i][j] = (c[i][j] + a[i][k] * b[k][j]) % MOD
            }
        }
    }

    return c 
}
