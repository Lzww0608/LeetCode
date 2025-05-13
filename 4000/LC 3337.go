const MOD int = 1_000_000_007

type matrix [][]int

func newMatrix(n, m int) matrix {
    a := make(matrix, n)
    for i := range a {
        a[i] = make([]int, m)
    }

    return a
}

func (a matrix) mul(b matrix) matrix {
    c := newMatrix(len(a), len(b[0]))
    for i, row := range a {
        for k, x := range row {
            if x == 0 {
                continue
            }
            for j, y := range b[k] {
                c[i][j] = (c[i][j] + x * y) % MOD
            }
        }
    }

    return c
}

func (a matrix) powMul(n int, f matrix) matrix {
    res := f 
    for n > 0 {
        if n & 1 == 1 {
            res = a.mul(res)
        }
        a = a.mul(a)
        n >>= 1
    }

    return res
}

const N int = 26

func lengthAfterTransformations(s string, t int, nums []int) (ans int) {
    f := newMatrix(26, 1)
    for i := range f{
        f[i][0] = 1
    }

    m := newMatrix(N, N)
    for i, c := range nums {
        for j := i + 1; j <= i + c; j++ {
            m[i][j % N] = 1
        }
    }

    mt := m.powMul(t, f)
    cnt := [N]int{}
    for _, c := range s {
        cnt[c-'a']++
    }
    for i, row := range mt {
        ans = (ans + row[0] * cnt[i]) % MOD
    }

    return
}