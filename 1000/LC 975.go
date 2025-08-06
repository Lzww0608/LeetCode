func oddEvenJumps(arr []int) int {
    n := len(arr)
    f := make([][2]bool, n)
    f[n-1][0], f[n-1][1] = true, true 

    type pair struct {
        x, i int
    }
    a := make([]pair, n)
    for i, x := range arr {
        a[i] = pair{x, i}
    }

    solve := func() []int {
        next := make([]int, n)
        for i := range next {
            next[i] = -1
        }
        st := []int{}
        for _, v := range a {
            for len(st) > 0 && st[len(st)-1] < v.i {
                next[st[len(st)-1]] = v.i
                st = st[:len(st)-1]
            }

            st = append(st, v.i)
        }

        return next
    }

    sort.Slice(a, func(i, j int) bool {
        if a[i].x == a[j].x {
            return a[i].i < a[j].i
        }
        return a[i].x < a[j].x
    })
    odd := solve()

    sort.Slice(a, func(i, j int) bool {
        if a[i].x == a[j].x {
            return a[i].i < a[j].i
        }
        return a[i].x > a[j].x
    })
    even := solve()

    ans := 1
    for i := n - 2; i >= 0; i-- {
        if odd[i] != -1 {
            f[i][0] = f[odd[i]][1]
        }
        if even[i] != -1 {
            f[i][1] = f[even[i]][0]
        }

        if f[i][0] {
            ans++
        }
    }

    return ans
}