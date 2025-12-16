func minDeletions(str string, queries [][]int) []int {
    n := len(str)
    s := []byte(str)
    f := make([]int, n)
    update := func(i, val int) {
        for ; i < n; i += i & (-i) {
            f[i] += val
        }
    }

    query := func(i int) (res int) {
        for ; i > 0; i -= i & (-i) {
            res += f[i]
        }

        return 
    }

    pre := func(l, r int) int {
        if l > r {
            return 0
        }

        return query(r) - query(l - 1) 
    }

    for i := 1; i < n; i++ {
        if s[i] == s[i - 1] {
            update(i, 1)
        }
    }

    ans := []int{}
    for _, q := range queries {
        if q[0] == 2 {
            ans = append(ans, pre(q[1] + 1, q[2]))
        } else {
            i := q[1]
            if i > 0 {
                if s[i] != s[i - 1] {
                    update(i, 1)
                } else {
                    update(i, -1)
                }
            }

            if i < n - 1 {
                if s[i] != s[i + 1] {
                    update(i + 1, 1)
                } else {
                    update(i + 1, -1)
                }
            }

            if s[i] == 'A' {
                s[i] = 'B'
            } else {
                s[i] = 'A'
            }
        }
    }

    return ans
}