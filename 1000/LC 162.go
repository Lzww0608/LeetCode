func findPeakElement(nums []int) int {
    l, r := -1, len(nums) - 1
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if nums[mid] > nums[mid + 1] {
            r = mid
        } else {
            l = mid
        }
    }

    return r
}