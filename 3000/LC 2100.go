func goodDaysToRobBank(security []int, time int) (ans []int) {
    st := []int{}
    n := len(security)
    suf := make([]int, n)
    
    for i := n - 1; i >= 0; i-- {
        x := security[i]
        if len(st) > 0 && x > st[len(st) - 1] {
            st = st[:0]
        }
        st = append(st, x)
        suf[i] = len(st)
    }

    st = st[:0]
    for i, x := range security {
        if len(st) > 0 && x > st[len(st) - 1] {
            st = st[:0]
        }
        st = append(st, x)
        if len(st) >= time + 1 && suf[i] >= time + 1 {
            ans = append(ans, i)
        }
    }

    return
}