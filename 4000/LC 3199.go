
import "math/bits"
func tripletCount(a []int, b []int, c []int) int {
    x := count(a)
    y := count(b)
    z := count(c)
    ans := len(a) * len(b) * len(c)
    return ans - x * y * z - x * (len(b) - y) * (len(c) - z) - y * (len(a) - x) * (len(c) - z) - z * (len(a) - x) * (len(b) - y)
}

func count(a []int) int {
    ans := 0
    for _, x := range a {
        if bits.OnesCount(uint(x)) & 1 == 1 {
            ans++
        }
    }

    return ans
}