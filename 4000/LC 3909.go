func compareBitonicSums(nums []int) int {
    sum, i := 0, 0
    n := len(nums)
    for i = 1; i < n; i++ {
        if nums[i] > nums[i - 1] {
            sum += nums[i - 1]
        } else {
            break
        }
    }

    for i < n {
        sum -= nums[i]
        i++
    }

    if sum > 0 {
        return 0
    } else if sum < 0 {
        return 1
    }
    return -1
}