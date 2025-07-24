
import "slices"
func maxSum(nums []int) (ans int) {
    mx := slices.Max(nums)
    if mx <= 0 {
        return mx
    }

    vis := make(map[int]bool)
    for _, x := range nums {
        if x >= 0 && !vis[x] {
            ans += x
            vis[x] = true
        }
    }

    return
}