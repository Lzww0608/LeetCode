func findSubarrays(nums []int) bool {
    n := len(nums)
    m := make(map[int]int)
    for i := 1; i < n - 1; i++ {
        m[nums[i] + nums[i + 1]]++
    }

    for i := 0; i < n - 2; i++ {
        if m[nums[i] + nums[i + 1]] > 0 {
            return true
        }
        m[nums[i + 1] + nums[i + 2]]--
    }

    return false
}