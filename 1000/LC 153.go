func findMin(nums []int) int {
    n := len(nums)

    l, r := 0, n 
    for l < r {
        mid := l + ((r - l) >> 1)
        if nums[mid] > nums[n-1] {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return nums[r]
}