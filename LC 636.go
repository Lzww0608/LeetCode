type pair struct {
    x, y int
}
func exclusiveTime(n int, logs []string) []int {
    ans := make([]int, n)
    st := []pair{}
    for _, s := range logs {
        id, i, m := 0, 0, len(s)
        for s[i] != ':' {
            id = id * 10 + int(s[i] - '0')
            i++
        }
        f := true
        i++
        if s[i] == 'e' {
            f = false
        }
        for s[i] != ':' {
            i++
        }
        i++
        time := 0
        for i < m {
            time = time * 10 + int(s[i] - '0')
            i++
        }
        if f {
            if len(st) > 0 {
                ans[st[len(st)-1].x] += time - st[len(st)-1].y
                st[len(st)-1].y = time
            } 
            st = append(st, pair{id, time})
        } else {
            t := st[len(st)-1]
            st = st[:len(st)-1]
            ans[t.x] += time - t.y + 1
            if len(st) != 0 {
                st[len(st)-1].y = time + 1
            }
        }
    }
    return ans
}