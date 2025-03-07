func validSubarrays(nums []int) (ans int) {
    n := len(nums)
    st := []int{n}
    for i := n - 1; i >= 0; i-- {
        x := nums[i]
        for len(st) > 1 && nums[st[len(st)-1]] >= x {
            st = st[:len(st)-1]
        }
        ans += st[len(st)-1] - i 
        st = append(st, i)
    }

    return 
}