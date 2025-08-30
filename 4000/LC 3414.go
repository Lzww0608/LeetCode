func maximumWeight(intervals [][]int) []int {
    n := len(intervals)
    type tuple struct {l, r, w, i int}
    a := make([]tuple, n)
    for i, v := range intervals {
        a[i] = tuple{v[0], v[1], v[2], i}
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i].r < a[j].r
    })

    type pair struct {
        mx int 
        a []int
    }
    f := make([][5]pair, n + 1)
    for i, v := range a {
        k := sort.Search(i, func(k int) bool {return a[k].r >= v.l})
        for j := 1; j < 5; j++ {
            s1 := f[i][j].mx
            s2 := f[k][j - 1].mx + v.w
            if s1 > s2 {
                f[i + 1][j] = f[i][j]
                continue
            }

            cur := slices.Clone(f[k][j - 1].a)
            cur = append(cur, v.i)
            sort.Ints(cur)
            if s1 == s2 && slices.Compare(f[i][j].a, cur) < 0 {
                cur = f[i][j].a
            }
            f[i + 1][j] = pair{s2, cur}
        }
    }

    return f[n][4].a
}