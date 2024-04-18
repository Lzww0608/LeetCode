func minKBitFlips(nums []int, k int) int {
    ans, d := 0, 0

    for i := range nums {
        if i >= k && nums[i-k] > 1 {
            d--  
        }

        if d & 1 == nums[i] {
            if i + k > len(nums) {
                return -1
            }
            d++
            ans++
            nums[i] += 2
        }
    }
    

    return ans
}