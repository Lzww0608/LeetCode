func countSubarrays(nums []int, k int) int64 {
    ans := 0
    m := make(map[int]int)
    for i, x := range nums {
        m[x]++
        for j := i - 1; j >= 0 && nums[j] & x != nums[j]; j-- {
            m[nums[j]]--
            nums[j] &=x 
            m[nums[j]]++
        }
        ans += m[k]
    }

    return int64(ans)
}