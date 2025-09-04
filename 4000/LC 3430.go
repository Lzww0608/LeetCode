
import "math"
func minMaxSubarraySum(nums []int, k int) int64 {
    solve := func() (res int) {
        st := []int{-1}
        for r, x := range nums {
            for len(st) > 1 && nums[st[len(st) - 1]] >= x {
                i := st[len(st) - 1]
                st = st[:len(st) - 1]
                l := st[len(st) - 1]
                if r - l - 1 <= k {
                    res += nums[i] * (i - l) * (r - i)
                } else {
                    l = max(l, i - k)
                    t := min(r, i + k)
                    a := (t - i) * (i - (t - k))
                    b := (l + t + k - i * 2 + 1) * (t - l - k) / 2
                    res += nums[i] * (a + b)
                }
            }
            st = append(st, r)
        }
        return
    }
    nums = append(nums, math.MinInt32)
    ans := solve()
    for i := range nums[:len(nums) - 1] {
        nums[i] = -nums[i]
    }
    ans -= solve()

    return int64(ans)
}