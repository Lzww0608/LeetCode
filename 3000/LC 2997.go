
import "math/bits"
func minOperations(nums []int, k int) int {
    for _, x := range nums {
        k ^= x
    }

    return bits.OnesCount(uint(k))
}