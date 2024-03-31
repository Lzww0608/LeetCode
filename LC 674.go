func findLengthOfLCIS(nums []int) (ans int) {
    n := len(nums)
    cnt := 1
    ans = 1
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            cnt++
        } else {
            cnt = 1
        }
        ans = max(ans, cnt)
    }

    return 
}