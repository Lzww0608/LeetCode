func countSubarrays(nums []int) (ans int) {
    for i := 0; i + 2 < len(nums); i++ {
        if (nums[i] + nums[i+2]) * 2 == nums[i+1] {
            ans++
        }
    }

    return
}