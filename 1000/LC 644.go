const EPS float64 = 1e-5
func findMaxAverage(nums []int, k int) float64 {
    n := len(nums)

    check := func(limit float64) bool {
        mn := math.MaxFloat32 
        sum := make([]float64, n + 1)
        for i, x := range nums {
            sum[i+1] = sum[i] + float64(x) - limit
            if i >= k - 1 {
                mn = min(mn, sum[i-k+1])
            }
            if sum[i+1] - mn >= 0 {
                return true
            }
        }
        return false
    }

    l, r := float64(slices.Min(nums) - 1), float64(slices.Max(nums) + 1)
    for l + EPS < r {
        mid := l + ((r - l) / 2)
        if check(mid) {
            l = mid 
        } else {
            r = mid
        }
    }

    return l
}