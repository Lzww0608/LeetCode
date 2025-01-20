func maximumLengthOfRanges(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    st := []int{-1}

    for i, x := range nums {
        for len(st) > 1 && x > nums[st[len(st)-1]] {
            cur := st[len(st)-1]
            st = st[:len(st)-1]
            ans[cur] = i - st[len(st)-1] - 1
        }
        st = append(st, i)
    }
    
    pre := -1
    for _, i := range st[1:] {
        ans[i] = n - pre - 1
        pre = i 
    }

    return ans
}