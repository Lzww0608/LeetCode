
import "slices"
func maximumLength(nums []int, k int) (ans int) {
    for i := 0; i <= k; i++ {
        f := make([]int, k)
        for _, x := range nums {
            y := x % k 
            f[y] = max(f[y], f[(i - y + k) % k] + 1)
        }

        ans =  max(ans, slices.Max(f))
    }

    return
}