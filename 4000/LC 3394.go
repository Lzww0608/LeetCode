type pair struct {
    l, r int
}
func checkValidCuts(n int, rectangles [][]int) bool {
    m := len(rectangles)
    
    a := make([]pair, m)
    b := make([]pair, m)

    for i, v := range rectangles {
        a[i] = pair{v[0], v[2]}
        b[i] = pair{v[1], v[3]}
    }

    sort.Slice(a, func(i, j int) bool {
        return a[i].l < a[j].l
    })
    sort.Slice(b, func(i, j int) bool {
        return b[i].l < b[j].l
    })

    return check(a) || check(b)
}

func check(a []pair) bool {
    cnt := 0

    r := -1
    for _, v := range a {
        if v.l >= r {
            cnt++
        }
        r = max(r, v.r)
    }

    return cnt >= 3
}