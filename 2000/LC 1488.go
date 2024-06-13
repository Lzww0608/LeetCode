func avoidFlood(rains []int) []int {
    n := len(rains)
    ans := make([]int, n)
    m := map[int]int{}
    st := []int{}
    for i := range ans {
        ans[i] = 1
    }

    for i, x := range rains {
        if x == 0 {
            st = append(st, i)
        } else {
            ans[i] = -1
            if v, ok := m[x]; ok {
                pos := sort.SearchInts(st, v)
                if pos == len(st) {
                    return []int{}
                }
                ans[st[pos]] = x
                copy(st[pos:len(st)-1], st[pos+1:len(st)])
                st = st[:len(st)-1]
            }
        }
        m[x] = i
    }

    return ans
}
