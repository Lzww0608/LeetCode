
import "math"
func partitionArray(nums []int, k int) int {
    ans := 1
    sort.Ints(nums)
    mn, mx := math.MaxInt32, math.MinInt32

    for _, x := range nums {
        mn = min(mn, x)
        mx = max(mx, x)
        if mx - mn > k {
            ans++
            mn, mx = x, x
        }
    }
    
    return ans
}