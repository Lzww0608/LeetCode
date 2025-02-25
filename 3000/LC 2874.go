
import "math"
func maximumTripletValue(nums []int) int64 {

    x, mx := 0, math.MinInt32
    ans := 0
    for _, y := range nums {
        ans = max(ans, x * y)
        x = max(x, mx - y)
        mx = max(mx, y)
    }
    if mx < 0 {
        return 0
    }
    return int64(ans)
}

