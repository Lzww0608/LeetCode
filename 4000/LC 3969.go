func countValidSubarrays(nums []int, x int) (ans int) {
    n := len(nums)
    pre := make([]int, n + 1)
    for i, t := range nums {
        pre[i + 1] = pre[i] + t
    }

    for low, high := x, x + 1; low <= pre[n]; low, high = low * 10, high * 10 {
        cnt := [10]int{}
        i, j := 0, 0
        for _, t := range pre {
            for pre[i] <= t - high {
                cnt[pre[i] % 10]--
                i++
            }
            for pre[j] <= t - low {
                cnt[pre[j] % 10]++
                j++
            }

            ans += cnt[(t - x + 10) % 10]
        }
    }

    return
}