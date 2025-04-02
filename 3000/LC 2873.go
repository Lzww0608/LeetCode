
import "math"
func maximumTripletValue(nums []int) int64 {
    ans := 0
    x, mx := 0, math.MinInt32
    for _, y := range nums {
        ans = max(ans, x * y)
        x = max(x, mx - y)
        mx = max(mx, y)
    }

    return int64(ans)
}