func findMin(nums []int) int {
    n := len(nums)

    l, r := 0, n - 1 
    for l < r {
        mid := l + ((r - l) >> 1)
        if nums[mid] < nums[r] {
            r = mid 
        } else if nums[mid] > nums[r] {
            l = mid + 1
        } else {
            r--
        }
    }

    return nums[r]
}