func maximumBeauty(nums []int, k int) (ans int) {
    sort.Ints(nums)
    i, j := 0, 0
    n := len(nums)
    for j < n {
        for j < n && nums[j] - nums[i] <= k * 2 {
            j++
        }
        ans = max(ans, j - i)
        i++
    }

    return 
}