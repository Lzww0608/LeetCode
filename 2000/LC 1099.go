func twoSumLessThanK(nums []int, k int) (ans int) {
    sort.Ints(nums)
    n := len(nums)
    l, r := 0, n - 1
    for l < r {
        x := nums[l] + nums[r]
        if x >= k {
            r--
        } else {
            ans = max(ans, x)
            l++
        }
    }

    if ans == 0 {
        return -1
    }

    return
}