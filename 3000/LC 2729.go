const N int = 9
func isFascinating(n int) bool {
    a, b, c := n, n * 2, n * 3 
    mask := 0
    for _, x := range []int{a, b, c} {
        for x > 0 {
            y := x % 10 - 1
            if y == -1 || mask & (1 << y) != 0 {
                return false
            }
            mask |= 1 << y
            x /= 10
        }
    }

    return mask == (1 << N) - 1
}