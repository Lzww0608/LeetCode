func minDeletions(nums []int) (ans int) {
    n := len(nums)
    for i := 1; i < n; i++ {
        if nums[i] ^ nums[i - 1] == 0 {
            ans++
        }
    }

    return 
}