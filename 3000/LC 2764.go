func isPreorder(nodes [][]int) bool {
    st := []int{-1}
    for _, v := range nodes {
        for len(st) > 0 && st[len(st)-1] != v[1] {
            st = st[:len(st)-1]
        }

        if len(st) == 0 {
            return false
        }

        st = append(st, v[0])
    }

    return true
}