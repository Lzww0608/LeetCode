func countSubarrays(nums []int, minK int, maxK int) int64 {
    ans := 0
    l := 0
    for r, x := range nums {
        if x < minK || x > maxK {
            l = r + 1
        }

        ans += r - l + 1
    }

    return int64(ans)
}