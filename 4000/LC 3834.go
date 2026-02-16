func mergeAdjacent(nums []int) []int64 {
    st := []int64{}
    for _, x := range nums {
        for len(st) > 0 && int(st[len(st) - 1]) == x {
            x += int(st[len(st) - 1])
            st = st[:len(st) - 1]
        }
        st = append(st, int64(x))

    }  

    return st
}