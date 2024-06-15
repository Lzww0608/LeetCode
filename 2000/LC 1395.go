import (
    "sort"
)

func numTeams(rating []int) (ans int) {
    n := len(rating)

    update := func(a []int, idx, val int) {
        for i := idx; i <= len(a); i += i & (-i) {
            a[i-1] += val
        }
    }

    query := func(a []int, idx int) (res int) {
        for i := idx; i > 0; i -= i & (-i) {
            res += a[i-1]
        }
        return res
    }

    tmp := make([]int, n)
    copy(tmp, rating)
    sort.Ints(tmp)
    m := map[int]int{}
    for i, x := range tmp {
        m[x] = i + 1
    }

    tr1 := make([]int, n+1)
    tr2 := make([]int, n+1)
    for _, x := range rating {
        update(tr1, m[x], 1)
    }

    for _, x := range rating {
        update(tr1, m[x], -1)
        ans += query(tr2, m[x]-1) * (query(tr1, n) - query(tr1, m[x]))
        ans += query(tr1, m[x]-1) * (query(tr2, n) - query(tr2, m[x]))
        update(tr2, m[x], 1)
    }

    return
}
