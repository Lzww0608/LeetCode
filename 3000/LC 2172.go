
import "math/bits"
func maximumANDSum(nums []int, numSlots int) (ans int) {
    f := make([]int, 1 << (numSlots << 1))
    for i, x := range f {
        c := bits.OnesCount(uint(i))
        if c >= len(nums) {
            continue
        }

        for j := 0; j < numSlots * 2; j++ {
            if (i >> j) & 1 == 0 {
                s := i | (1 << j)
                f[s] = max(f[s], x + (j / 2 + 1) & nums[c])
                ans = max(ans, f[s])
            }
        }
    }

    return
}
