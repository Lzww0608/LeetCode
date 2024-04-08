func searchRange(nums []int, target int) []int {

    l, r := 0, len(nums)
    for l < r {
        mid := l + ((r - l) >> 1)
        if nums[mid] >= target {
            r = mid 
        } else {
            l = mid + 1
        }
    }
    left := l

    l, r = 0, len(nums)
    for l < r {
        mid := l + ((r - l) >> 1)
        if nums[mid] > target {
            r = mid 
        } else {
            l = mid + 1
        }
    }
    right := l - 1

    if left <= right {
        return []int{left, right}
    } 
    return []int{-1, -1}
}