func getMaxLen(nums []int) (ans int) {
    n := len(nums)

    a := []int{-1}
    for i := 0; i <= n; i++ {
        if i == n || nums[i] == 0 {
            a = append(a, i)
            m := len(a)
            if m & 1 == 0 {
                ans = max(ans, a[m - 1] - a[0] - 1)
            } else {
                ans = max(ans, a[m - 1] - a[1] - 1, a[m - 2] - a[0] - 1)
            }

            a = []int{i}
        } else if nums[i] < 0 {
            a = append(a, i)
        } 
    }

    return 
}