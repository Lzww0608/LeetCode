func find132pattern(nums []int) bool {
    st := []int{}
    mx := math.MinInt32

    for i := len(nums) - 1; i >= 0; i-- {
        if mx > nums[i] {
            return true
        }

        for len(st) > 0 && nums[i] > st[len(st)-1] {
            mx = max(mx, st[len(st)-1])
            st = st[:len(st)-1]
        }

        st = append(st, nums[i])
    }

    return false
}
