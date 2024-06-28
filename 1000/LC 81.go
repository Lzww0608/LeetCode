func search(nums []int, target int) bool {
    n := len(nums)
    if n == 1 {
        if nums[0] == target {
            return true
        }
        return false
    }
    l, r := 0, n - 1

    for l <= r {
        mid := l + ((r - l) >> 1)
        if nums[mid] == target {
            return true
        } else if nums[mid] == nums[l] {
            l++
        } else if nums[mid] > nums[l] {
            if nums[l] <= target && target < nums[mid] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else {
            if target <= nums[r] && nums[mid] < target {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return false
}
