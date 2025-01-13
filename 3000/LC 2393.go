func countSubarrays(nums []int) int64 {
    ans := 0
    cnt := 0
    for i := range nums {
        if i > 0 && nums[i] > nums[i-1] {
            cnt++
        } else {
            cnt = 1
        }
        ans += cnt
    }

    return int64(ans)
}