func swapNumbers(a []int) []int {
    a[0] ^= a[1]
    a[1] ^= a[0]
    a[0] ^= a[1]
    return a
}