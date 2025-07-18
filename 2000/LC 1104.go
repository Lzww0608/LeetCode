import (
	"slices"
)
func pathInZigZagTree(label int) (ans []int) {
    ans = append(ans, label)
    for label > 1 {
        label = label / 2
        if label != 1 {
            a := nextPowerof2(label)
            b := a >> 1 
            label = a - label + b - 1
        }
        
        ans = append(ans, label)
    }
    slices.Reverse(ans)
    return
}

func nextPowerof2(x int) int {
    if x == 1 {
        return 1
    }
    x |= x >> 1
    x |= x >> 2
    x |= x >> 4
    x |= x >> 8
    x |= x >> 16

    return x + 1
}