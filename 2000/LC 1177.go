func canMakePaliQueries(s string, queries [][]int) []bool {
    n := len(queries)
    ans := make([]bool, n)
    pre := make([]int, len(s) + 1)
    for i := range s {
        pre[i+1] = pre[i] ^ (1 << int(s[i] - 'a'))
    }

    check := func(l, r int) (res int) {
        x := pre[r+1] ^ pre[l]
        for x > 0 {
            if x & 1 == 1 {
                res++
            }
            x >>= 1
        }
        return res / 2
    }
    for i, query := range queries {
        l, r, k := query[0], query[1], query[2]
        if check(l, r) <= k {
            ans[i] = true
        }
    }

    return ans
}
