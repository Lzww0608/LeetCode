func minOperationsToMakeMedianK(nums []int, k int) int64 {
    n := len(nums)
    sort.Ints(nums)
    ans := 0

    for i, x := range nums {
        if i < n / 2 {
            if x > k {
                ans += x - k
            }
        } else if i > n / 2 {
            if x < k {
                ans += k - x
            }
        } else {
            ans += max(x - k, k - x)
        }
    }

    return int64(ans)
}