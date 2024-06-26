func maximumCount(nums []int) int {
    n := len(nums)
    l, r := 0, n

    for l < r {
        mid := l + ((r - l) >> 1)
        if nums[mid] >= 0 {
            r = mid
        } else {
            l = mid + 1
        }
    }
    neg := l

    l, r = 0, n 
    for l < r {
        mid := l + ((r - l) >> 1)
        if nums[mid] > 0 {
            r = mid
        } else {
            l = mid + 1
        }
    }
    pos := n - l

    return max(neg, pos)
}
