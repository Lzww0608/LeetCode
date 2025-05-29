func check(nums []int) bool {
    cnt := 0
    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i] < nums[i-1] {
            cnt++
        }
    }

    return cnt == 0 || cnt == 1 && nums[0] >= nums[n-1]
}