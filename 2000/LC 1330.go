
import "math"
func maxValueAfterReverse(nums []int) (ans int) {
    n := len(nums)
    for i := 1; i < n; i++ {
        ans += abs(nums[i] - nums[i-1])
    }

    mx, mn := math.MinInt32, math.MaxInt32
    for i := 1; i < n; i++ {
        mx = max(mx, min(nums[i], nums[i-1]))
        mn = min(mn, max(nums[i], nums[i-1]))
    }
    d := max(0, 2 * (mx - mn))

    for i := 0; i < n - 1; i++ {
        d = max(d, abs(nums[i+1] - nums[0]) - abs(nums[i] - nums[i+1]))
        d = max(d, abs(nums[n-1] - nums[i]) - abs(nums[i] - nums[i+1]))
    }

    return ans + d
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}