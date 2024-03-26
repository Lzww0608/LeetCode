func removeDuplicates(nums []int) int {
    l, r, n := 0, 1, len(nums)
    for r < n {
        if nums[r] != nums[l] {
            l++
            nums[l] = nums[r]
        }
        r++
    }

    return l + 1
}
