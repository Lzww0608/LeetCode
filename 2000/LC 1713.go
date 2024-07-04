func minOperations(target []int, arr []int) int {
    m := map[int]int{}
    n := len(target)
    for i, x := range target {
        m[x] = i
    }

    st := []int{math.MinInt}
    for _, x := range arr {
        if _, ok := m[x]; !ok {
            continue
        }
        x := m[x]
        if x > st[len(st)-1] {
            st = append(st, x)
            continue
        }
        pos := sort.SearchInts(st, x)
        st[pos] = x
    }

    return n - len(st) + 1
}