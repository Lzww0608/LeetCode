func countAlternatingSubarrays(nums []int) int64 {
    n := len(nums)
    var ans int64 = int64(n)
    var cnt int64 = 1
    for i := 1; i < n; i++ {
        if nums[i] != nums[i-1] {
            cnt++
        } else {
            ans += cnt * (cnt - 1) / 2
            cnt = 1
        }
    }
    ans += cnt * (cnt - 1) / 2
    return ans
}