type pair struct {
    l, r int 
}
type ST [][]int 

func NewST(a []pair) ST {
    n := len(a) - 1
    st := make(ST, n)
    m := bits.Len(uint(n))
    for i := range st {
        st[i] = make([]int, m)
        st[i][0] = a[i].r - a[i].l + a[i + 1].r - a[i + 1].l
    }

    for j := range m - 1 {
        for i := 0; i + (1 << (j + 1)) <= n; i++ {
            st[i][j + 1] = max(st[i][j], st[i + (1 << j)][j])
        }
    }

    return st
}

func (st ST)Query(l, r int) int {
    if l >= r {
        return 0
    }

    k := bits.Len(uint(r - l)) - 1
    return max(st[l][k], st[r - (1 << k)][k])
}

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
    n := len(s)
    a := []pair{{-1, -1}}
    p := make([]int, n)
    l, r := 0, 0
    sum := 0
    for i := range s {
        p[i] = len(a)
        
        if s[i] == '1' {
            sum++
        } else {
            r = i
        }

        if i == n - 1 || s[i] != s[i + 1] {
            if s[i] == '1' {
                l, r = i + 1, i + 1
            } else {
                a = append(a, pair{l, r + 1})
            }
        }
    }

    a = append(a, pair{n + 1, n + 1})

    solve := func(x, y int) int {
        if x > 0 && y > 0 {
            return x + y
        }
        return 0
    }

    st := NewST(a)
    ans := make([]int, len(queries))
    for k, q := range queries {
        l, r := q[0], q[1]

        i := p[l]
        if l > 0 && s[l] == '0' && s[l - 1] == '0' {
            i++
        }
        j := p[r] - 1
        if r + 1 < n && s[r] == '0' && s[r + 1] == '1' {
            j++
        }
        r++

        cur := 0
        if i <= j {
            cur = max(st.Query(i, j), solve(a[i - 1]. r - l, a[i].r - a[i].l), solve(r - a[j + 1].l, a[j].r - a[j].l))
        } else if i == j + 1 {
            cur = solve(a[i - 1].r - l, r - a[j + 1].l)
        }

        ans[k] = cur + sum
    }

    return ans
}