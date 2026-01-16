func centeredSubarrays(nums []int) (ans int) {
    for i := range nums {
        vis := make(map[int]bool)
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            vis[nums[j]] = true
            if vis[sum] {
                ans++
            }
        }
    }

    return 
}