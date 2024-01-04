func numSubarrayProductLessThanK(nums []int, k int) int {
    ans, n := 0, len(nums)
    l, r, prod := 0, 0, 1
    for r < n {
        if nums[r] >= k {
            l, r, prod = r + 1, r + 1, 1
            continue
        }
        prod *= nums[r]
        if prod < k {
            ans += r - l + 1
            r++
            continue
        }
        prod /= nums[r]
        prod /= nums[l]
        l++
    }  
    return ans
}