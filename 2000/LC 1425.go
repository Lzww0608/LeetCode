func constrainedSubsetSum(nums []int, k int) int {
    ans := nums[0]
    n := len(nums)
    f := make([]int, n)
    st := []int{}

    for i, x := range nums {
        for len(st) > 0 && st[0] < i - k {
            st = st[1:]
        }

        if len(st) == 0 {
            f[i] = x 
        } else {
            f[i] = max(f[st[0]] + x, x)
        }
        ans = max(ans, f[i])
        
        for len(st) > 0 && f[st[len(st)-1]] <= f[i] {
            st = st[:len(st)-1]
        }

        st = append(st, i)
    }

    return ans
}