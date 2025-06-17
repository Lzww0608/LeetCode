func canBeIncreasing(nums []int) bool {
    n := len(nums)
    cnt := 0
    for i := 1; i < n; i++ {
        if nums[i] <= nums[i-1] {
            cnt++
            if cnt > 1 {
                return false
            }
            if i == 1 || nums[i-2] < nums[i] || i == n - 1 || nums[i+1] > nums[i-1] {
                continue
            } else {
                return false
            }
        }
    }

    return true
}