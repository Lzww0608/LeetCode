func maxProduct(nums []int) int64 {
    sort.Ints(nums)
    n := len(nums)
    var ans int = 0
    if nums[n - 1] > 0 && nums[n - 2] > 0 {
        ans = nums[n - 1] * nums[n - 2] * 100_000
    }

    if nums[0] < 0 {
        if nums[1] < 0 {
            ans = max(ans, nums[0] * nums[1] * 100_000)
        } 
        if nums[n - 1] > 0 {
            ans = max(ans, nums[0] * nums[n - 1] * -100_000)
        }
    }

    return int64(ans)
}