
import (
	"math"
	"slices"
)
func maxSubarraySumCircular(nums []int) int {
    ans := math.MinInt
    sum := 0
    
    cur := 0
    for _, x := range nums {
        cur = max(cur + x, x)
        ans = max(ans, cur)
        sum += x
    }
    if slices.Max(nums) <= 0 {
        return ans
    }

    cur = 0
    for _, x := range nums {
        cur = min(cur + x, x)
        ans = max(ans, sum - cur)
    }
    
    return ans
}