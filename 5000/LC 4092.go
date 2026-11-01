func calculateScore(a int, b int) int {
    x := bits.OnesCount(uint(a))
    y := bits.OnesCount(uint(b))
    if x >= y {
        return x * y
    }

    return y / x
}