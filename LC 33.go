func search(nums []int, target int) int {
    n := len(nums)
    if n == 1 {
        if nums[0] == target {
            return 0
        }
        return -1
    }
    l, r := 0, n - 1

    for l <= r {
        mid := l + ((r - l) >> 1)
        if nums[mid] == target {
            return mid
        } else if nums[mid] >= nums[0] {
            if nums[0] <= target && target < nums[mid] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if target <= nums[n-1] && nums[mid] < target {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return -1
}