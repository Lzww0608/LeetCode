func search(nums []int, target int) int {
    n := len(nums)
    l, r := 0, n
    for l < r {
        mid := l + ((r - l) >> 1) 
        if nums[mid] > target {
            r = mid
        } else if nums[mid] < target {
            l = mid + 1
        } else {
            return mid
        }
    }

    return -1
}
