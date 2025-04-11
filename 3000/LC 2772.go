func checkArray(nums []int, k int) bool {
    if k == 1 {
        return true
    }
    n := len(nums)
    diff := make([]int, n + 1)
    diff[0] = nums[0];
    for i := 1; i < n; i++ {
        diff[i] = nums[i] - nums[i-1]
    }

    for i := 0; i < n; i++ {
        if diff[i] == 0 {
            continue
        }

        if diff[i] < 0 || i + k > n {
            return false
        }

        diff[i+k] += diff[i]
    }

    return true
}