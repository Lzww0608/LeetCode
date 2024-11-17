
import "math"
func minMoves(nums []int, k int) int {
    n := len(nums)
    p := make([]int, 0, n)
    for i, x := range nums {
        if x == 1 {
            p = append(p, i - len(p))
        }
    }

    sum := make([]int, len(p) + 1)
    for i, x := range p {
        sum[i+1] = sum[i] + x
    }

    ans := math.MaxInt32
    for i := 0; i < len(sum) - k; i++ {
        ans = min(ans, sum[i] + sum[i+k] - sum[i+k/2] * 2 - p[i+k/2] * (k & 1))
    }

    return ans
}