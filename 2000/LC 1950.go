func findMaximums(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    st := []int{-1}

    for i, x := range nums {
        for len(st) > 1 && x < nums[st[len(st)-1]] {
            j := st[len(st)-1]
            st = st[:len(st)-1]
            k := st[len(st)-1]
            ans[i-k-2] = max(ans[i-k-2], nums[j])
        }

        st = append(st, i)
    }

    for len(st) > 1 {
        r := st[len(st)-1]
        st = st[:len(st)-1]
        l := st[len(st)-1]
        ans[n-l-2] = max(ans[n-l-2], nums[r])
    }

    for i := n - 2; i >= 0; i-- {
        ans[i] = max(ans[i], ans[i+1])
    }

    return ans
}