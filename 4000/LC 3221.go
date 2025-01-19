func maxScore(nums []int) int64 {
    ans := 0
    st := []int{}
    for i, x := range nums {
        for len(st) > 0 && nums[st[len(st)-1]] <= x {
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }

    for len(st) > 1 {
        a, b := st[len(st)-2], st[len(st)-1]
        ans += nums[b] * (b - a)
        st = st[:len(st)-1]
    }
    ans += nums[st[0]] * (st[0] - 0)
    return int64(ans)
}