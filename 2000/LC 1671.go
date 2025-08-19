func minimumMountainRemovals(nums []int) int {
    n := len(nums)
    suf := make([]int, n)
    st := []int{}

    for i := n - 1; i >= 0; i-- {
        if len(st) > 0 && nums[i] <= st[len(st) - 1] {
            pos := sort.SearchInts(st, nums[i])
            st[pos] = nums[i]
            suf[i] = pos + 1
        } else {
            st = append(st, nums[i])
            suf[i] = len(st)
        }
    }

    st = st[:0]
    ans := 0
    for i, x := range nums {
        j := sort.SearchInts(st, x)
        if j == len(st) {
            st = append(st, x)
        } else {
            st[j] = x
        }

        pre := j + 1
        if pre >= 2 && suf[i] >= 2 {
            ans = max(ans, pre + suf[i] - 1)
        }
    }
    
    return n - ans
}