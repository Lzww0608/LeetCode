func minimumOperations(nums []int) int {
    n := len(nums)
    st := []int{-1}
    for _, x := range nums {
        if x >= st[len(st)-1] {
            st = append(st, x)
        } else {
            pos := sort.SearchInts(st, x + 1)
            st[pos] = x
        }
    }

    return n - len(st) + 1
}