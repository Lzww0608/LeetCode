func countPairs(nums []int, k int) (ans int) {
    n := len(nums)
    for i := range nums {
        for j := i + 1; j < n; j++ {
            if nums[i] == nums[j] && i * j % k == 0 {
                ans++
            }
        }
    }

    return 
}
