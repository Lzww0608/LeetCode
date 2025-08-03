func maxSubarrayLength(nums []int) (ans int) {
    n := len(nums)
    st := []int{}
    for i, x := range nums {
        if len(st) == 0 || nums[st[len(st)-1]] < x {
            st = append(st, i)
        }
    }

    for i := n - 1; i > 0; i-- {
        x := nums[i]
        for len(st) > 0 && nums[st[len(st)-1]] > x {
            j := st[len(st)-1]
            st = st[:len(st)-1]
            ans = max(ans, i - j + 1)
        }
    }

    return
}