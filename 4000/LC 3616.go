
import "math"
func totalReplacements(ranks []int) (ans int) {
    pre := math.MaxInt
    for _, x := range ranks {
        if x < pre {
            pre = x
            ans++
        }
    }

    return ans - 1
}