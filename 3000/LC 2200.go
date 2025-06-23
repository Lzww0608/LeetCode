func findKDistantIndices(nums []int, key int, k int) (ans []int) {
    n := len(nums)
    for l, r := 0, 0; r < n; r++ {
        if nums[r] != key {
            continue
        }
        l = max(l, r - k)
        for l < n && l <= r + k {
            ans = append(ans, l)
            l++
        }
    } 
    
    return
}