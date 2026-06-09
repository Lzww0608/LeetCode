func maxTotal(nums []int, s string) int64 {
    ans := 0
    sum, cnt, mn := 0, 0, math.MaxInt

    n := len(nums)
    for i := range n {
        if s[i] == '0' {
            sum, cnt, mn = 0, 0, math.MaxInt
            continue
        }

        sum += nums[i]
        mn = min(mn, nums[i])
        if i == n - 1 || s[i + 1] == '0' {
            if cnt < i {
                ans += max(sum, sum - mn + nums[i - cnt - 1])
            } else {
                ans += sum
            }
        }
        cnt++
    }

    return int64(ans)
}