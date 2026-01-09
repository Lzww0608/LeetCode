func minMergeCost(lists [][]int) int64 {
    n := len(lists)
    u := 1 << n
    s := make([][]int, u)

    for i := range n {
        j := 1 << i 
        for k := range j {
            s[j | k] = merge(lists[i], s[k])
        }
    }

    f := make([]int, u)
    for i := 3; i < u; i++ {
        if i & (i - 1) == 0 {
            continue
        }
        f[i] = math.MaxInt
        for j := i & (i - 1); j > 0; j = (j - 1) & i {
            k := i ^ j 
            x := s[j][(len(s[j]) - 1) / 2]
            y := s[k][(len(s[k]) - 1) / 2]
            f[i] = min(f[i], f[j] + f[k] + abs(x - y))
        }
        f[i] += len(s[i])
    }

    return int64(f[u - 1])
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}

func merge(a, b []int) []int {
    n, m := len(a), len(b)
    res := make([]int, 0, n + m)
    for i, j := 0, 0; i < n || j < m; {
        if j >= m || i < n && a[i] < b[j] {
            res = append(res, a[i])
            i++ 
        } else {
            res = append(res, b[j])
            j++
        }
    }

    return  res
}