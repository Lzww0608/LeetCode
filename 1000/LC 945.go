
import "slices"
func minIncrementForUnique(nums []int) (ans int) {
    mx := slices.Max(nums)

    cnt := make([]int, mx + 1)
    for _, x := range nums {
        cnt[x]++
    }

    l := -1
    for i := 0; i <= mx; i++ {
        x := cnt[i]
        if x > 1 {
            cur := x - 1
            if i + 1 <= mx {
                ans += cur 
                cnt[i + 1] += cur
            } else {
                l = cur
            }
        }
    }

    return ans + l * (l + 1) / 2
}