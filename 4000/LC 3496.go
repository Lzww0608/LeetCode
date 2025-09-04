
import "math"
func maxScore(nums []int) int {
    sum := 0
    mn := math.MaxInt
    n := len(nums)
    for i, x := range nums {
        sum += x
        if n & 1 == 1 {
            mn =  min(mn, x)
        } else if i + 1 < len(nums) {
            mn = min(mn, nums[i] + nums[i + 1])
        }
    }

    return sum - mn
}