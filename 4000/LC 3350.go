func maxIncreasingSubarrays(nums []int) (ans int) {
    pre, cnt := 0, 0
    for i, x := range nums {
        if i > 0 && x > nums[i-1] {
            cnt++
        } else {
            ans = max(ans, min(pre, cnt), cnt / 2)
            pre = cnt
            cnt = 1
        }
    }
    ans = max(ans, min(pre, cnt), cnt / 2)
    return
}