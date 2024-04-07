func repeatedNTimes(nums []int) int {
    n := len(nums)

    for {
        i, j := rand.Intn(n), rand.Intn(n)
        if i != j && nums[i] == nums[j] {
            return nums[i]
        }
    }

    return -1
}