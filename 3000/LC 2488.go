func countSubarrays(nums []int, k int) int {
    n := len(nums)
    pos := 0
    for nums[pos] != k {
        pos++
    }
    cnt := make([]int, n * 2)
    cnt[n] = 1

    for i, x := pos - 1, n; i >= 0; i-- {
        if nums[i] < k {
            x++
        } else {
            x--
        }

        cnt[x]++
    }

    ans := cnt[n] + cnt[n-1]
    for i, x := pos + 1, n; i < n; i++ {
        if nums[i] < k {
            x--
        } else {
            x++
        }
        ans += cnt[x] + cnt[x-1]
    }

    return ans
}