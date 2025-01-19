func binarySearchableNumbers(nums []int) int {
    st := []int{}
    pre := nums[0]
    for i, x := range nums {
        for len(st) > 0 && x < nums[st[len(st)-1]] {
            st = st[:len(st)-1]
        }
        if x >= pre {
            st = append(st, i)
        }
        pre = max(pre, x)
    }

    return len(st)
}