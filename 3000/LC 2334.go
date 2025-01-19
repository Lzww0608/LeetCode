func validSubarraySize(nums []int, threshold int) int {
    st := []int{-1}
    for i, x := range nums {
        for len(st) > 1 && x < nums[st[len(st)-1]] {
            cur := st[len(st)-1]
            st = st[:len(st)-1]
            k := i - st[len(st)-1] - 1
            if threshold / k < nums[cur] {
                return k
            }
        }

        st = append(st, i)
    }

    pre := -1
    n := len(nums)
    for _, i := range st[1:] {
        if threshold / (n - pre - 1) < nums[i] {
            return n - pre - 1
        }
        pre = i 
    }

    return -1
}