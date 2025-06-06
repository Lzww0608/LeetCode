func findBuildings(heights []int) []int {
    st := []int{}
    for i, x := range heights {
        for len(st) > 0 && heights[st[len(st)-1]] <= x {
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }

    return st
}