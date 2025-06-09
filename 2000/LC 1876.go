
import "math/bits"
func countGoodSubstrings(s string) (ans int) {
    for i := 0; i + 2 < len(s); i++ {
        mask := 0
        for j := i; j <= i + 2; j++ {
            mask |= 1 << int(s[j] - 'a')
        }
        if bits.OnesCount(uint(mask)) == 3 {
            ans++
        }
    }

    return 
}