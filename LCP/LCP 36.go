func maxGroupNumber(tiles []int) int {
    cnt := make(map[int]int)
    for _, x := range tiles {
        cnt[x]++
    }

    a := make([]int, 0, len(cnt))
    for k := range cnt {
        a = append(a, k)
    }
    sort.Ints(a)

    f := [3][3]int{}
    for i := range 3 {
        for j := range 3 {
            f[i][j] = math.MinInt32
        }
    }
    f[0][0] = 0

    trans := func(c int) {
        g := [3][3]int{}
        for i := range 3 {
            for j := range 3 {
                g[i][j] = math.MinInt32
            }
        }

        for i := range 3 {
            for j := range 3 {
                if f[i][j] <= math.MinInt32 {
                    continue
                }
                for k := range 3 {
                    if i + j + k > c {
                        continue
                    }
                    g[j][k] = max(g[j][k], f[i][j] + (c - i - j - k) / 3 + k) 
                }
            }
        }

        f = g
    }

    last := 0
    for _, x := range a {
        c := cnt[x]
        gap := x - last - 1
        if gap >= 2 {
            trans(0)
        }
        if gap >= 1 {
            trans(0)
        }

        trans(c)
        last = x
    }

    trans(0)
    trans(0)

    return f[0][0]
}