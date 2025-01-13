func findScore(nums []int) int64 {
    n := len(nums)
    ans := 0
    for i := 0; i < n; i++ {
        j := i + 1
        for j < n && nums[j] < nums[j-1] {
            j++
        }

        for k := j - 1 ; k >= i; k -= 2 {
            ans += nums[k]
        }
        i = j
    }

    return int64(ans)
}