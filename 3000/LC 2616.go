func minimizeMax(nums []int, p int) int {
    sort.Ints(nums)
    n := len(nums)

    check := func(k int) bool {
        cnt := 0
        for i := 0; i < n - 1; i++ {
            if nums[i+1] - nums[i] <= k {
                cnt++
                i++
            } 
        }
        return cnt >= p
    }

    l, r := -1, nums[n-1] - nums[0]
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid 
        } else {
            l = mid
        }
    }

    return r
}
