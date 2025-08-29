
import "math"
func maxSubarraySum(nums []int) int64 {
    m := make(map[int]int)
    ans := math.MinInt
    all, non := 0, 0

    pre := 0
    for _, x := range nums {
        pre += x 
        ans = max(ans, pre - all)

        if x < 0 {
            m[x] = min(m[x], non) + x
            all = min(all, m[x])
            non = min(non, pre)
        }
    }

    return int64(ans)
}