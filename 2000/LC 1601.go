func maximumRequests(n int, requests [][]int) (ans int) {
    for i := 1; i < (1 << len(requests)); i++ {
        cnt := 0
        for j := i; j > 0; j >>= 1 {
            if j & 1 == 1 {
                cnt++
            }
        }
        if cnt <= ans {
            continue
        }
        m := map[int]int{}
        for j := range requests {
            if i & (1 << j) != 0 {
                a, b := requests[j][0], requests[j][1]
                m[a]++
                m[b]--
                if m[a] == 0 {
                    delete(m, a)
                }
                if m[b] == 0 {
                    delete(m, b)
                }
            }
        }
        if len(m) == 0 {
            ans = max(ans, cnt)
        }
    }

    return
}