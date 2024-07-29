func searchInsert(nums []int, target int) int {
    n := len(nums)
    l, r := 0, n 
    for l < r {
        mid := l + ((r - l) >> 1)
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            r = mid
        } else {
            l = mid + 1
        }
    }

    return r
}