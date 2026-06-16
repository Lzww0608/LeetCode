func maxRatings(units [][]int) int64 {
    m := len(units)
    if len(units[0]) == 1 {
        ans := 0
        for _, v := range units {
            ans += v[0]
        }
        return int64(ans)
    }

    sum := 0
    a := make([]pair, m)
    mn := math.MaxInt32
    for i, v := range units {
        a[i] = solve(v)
        sum += a[i].b
        mn = min(mn, a[i].a)
    }

    ans := 0
    for i := range m {
        ans = max(ans, sum - a[i].b + mn)
    }

    return int64(ans)
}

type pair struct {
    a, b int 
}

func solve(a []int) pair {
    t := pair{math.MaxInt32, math.MaxInt32}
    for _, x := range a {
        if x < t.a {
            t.a, t.b = x, t.a
        } else if x < t.b {
            t.b = x
        }
    }

    return t
}