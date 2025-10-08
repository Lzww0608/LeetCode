func maximumPossibleSize(nums []int) int {
    st := []int{}
    for _, x := range nums {
        if len(st) == 0 || x >= st[len(st) - 1] {
            st = append(st, x)
        }
    }

    return len(st)
}