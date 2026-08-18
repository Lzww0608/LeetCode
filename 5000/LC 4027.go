func elevatorRequests(n int, start int, requests [][]int) int64 {
    m := len(requests)
    f := make([][]int, 1 << m)
    for i := range f {
        f[i] = make([]int, m)
    }
    for i := range m {
        f[1 << i][i] = max(requests[i][0], abs(start - requests[i][1])) 
    }

    for s := 1; s < (1 << m); s++ {
        if s & (s - 1) == 0 {
            continue
        }

        for i, v := range requests {
            if s & (1 << i) == 0 {
                continue
            }

            cur := math.MaxInt
            mask := s ^ (1 << i)
            for j := range m {
                if mask & (1 << j) != 0 {
                    cur = min(cur, f[mask][j] + abs(v[1] - requests[j][1]))
                }
            }
            f[s][i] = max(cur, v[0])
        }
    } 

    return int64(slices.Min(f[(1 << m) - 1]))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}