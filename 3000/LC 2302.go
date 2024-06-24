func countSubarrays(nums []int, k int64) (ans int64) {
    var sum int64 = 0
    n := len(nums)
    l, r := 0, 0
    for r < n {
        sum += int64(nums[r])
        r++

        for l < r && sum * int64(r - l) >= k {
            sum -= int64(nums[l])
            l++
        }

        ans += int64(r - l)
    }

    return 
}
