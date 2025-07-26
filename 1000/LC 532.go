func findPairs(nums []int, k int) (ans int) {
    sort.Ints(nums)
    n := len(nums)
    for l, r := 0, 0; r < n; {
        if r == l {
            r++
        }
        for r < n && nums[r] - nums[l] < k {
            r++
        }
        if r < n && nums[r] - nums[l] == k {
            ans++
            for r < n && nums[r] - nums[l] == k {
                r++
            }
        }

        for x := nums[l]; l < n && nums[l] == x; l++ {}
    }

    return
}