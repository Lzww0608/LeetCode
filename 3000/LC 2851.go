const MOD int = 1_000_000_007

type matrix [][]int

func newMatrix(m, n, x int) matrix {
    a := make(matrix, m)
    for i := range a {
        a[i] = make([]int, n)
        for j := range a[i] {
            a[i][j] = x
        }
    }

    return a
}

func newIdentityMatrix(n int) matrix {
    a := make(matrix, n)
    for i := range a {
        a[i] = make([]int, n)
        a[i][i] = 1
    }

    return a
}

func (a matrix) mul(b matrix) matrix {
    c := newMatrix(len(a), len(b[0]), 0)
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(b[0]); j++ {
            for k := 0; k < len(a[0]); k++ {
                c[i][j] = (c[i][j] + a[i][k] * b[k][j]) % MOD
            }
        }
    }
    return c
}

func quickPow(a matrix, r int) matrix {
    ans := newIdentityMatrix(len(a))
    for r > 0 {
        if r & 1 == 1 {
            ans = ans.mul(a)
        }
        a = a.mul(a)
        r >>= 1
    }

    return ans
}

func numberOfWays(s string, t string, k int64) int {
    n := len(s)
    s = s + s 
    m := len(t)
    pi := make([]int, m)
    cnt := 0
    for i := 1; i < m; i++ {
        v := t[i]
        for cnt > 0 && t[cnt] != v {
            cnt = pi[cnt - 1]
        }
        if t[cnt] == v {
            cnt++
        }
        pi[i] = cnt
    }

    cnt = 0
    c := 0
    for i := range s {
        for cnt > 0 && t[cnt] != s[i] {
            cnt = pi[cnt - 1]
        }

        if t[cnt] == s[i] {
            cnt++
        }
        if cnt == m {
            if i - m + 1 < n {
                c++
            }
            cnt = pi[cnt - 1]
        }
    }

    ans := quickPow(matrix{{c - 1, c}, {n - c, n - 1 - c}}, int(k))
    if s == t + t {
        return ans[0][0]
    }
    return ans[0][1]
}