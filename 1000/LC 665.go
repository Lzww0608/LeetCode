func checkPossibility(nums []int) bool {
    n := len(nums)
    cnt := 0
    for i := 1; i < n; i++ {
        if nums[i] < nums[i-1] {
            if cnt >= 1 {
                return false
            }
            if i == 1 || nums[i] >= nums[i-2] || i == n - 1 || nums[i+1] >= nums[i-1] {
                cnt++
            } else {
                return false
            }
        }
    }

    return cnt <= 1
}