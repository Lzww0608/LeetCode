const MOD int = 1_000_000_007

func modInverse(x, mod int) int {
    return pow(x, mod-2, mod)
}

func pow(a, r, mod int) int {
    res := 1
    a = a % mod
    for r > 0 {
        if r & 1 == 1 {
            res = (res * a) % mod
        }
        a = (a * a) % mod
        r >>= 1
    }
    return res
}

func productQueries(n int, queries [][]int) []int {
    m := len(queries)
    ans := make([]int, m)
    a := []int{}
    
    for i := 1; i <= n; i <<= 1 {
        if n & i != 0 {
            a = append(a, i)
        }
    }

    pre := make([]int, len(a)+1)
    pre[0] = 1
    for i, x := range a {
        pre[i+1] = (pre[i] * x) % MOD
    }

    for i, q := range queries {
        l, r := q[0], q[1]
        ans[i] = (pre[r+1] * modInverse(pre[l], MOD)) % MOD
    }

    return ans
}
