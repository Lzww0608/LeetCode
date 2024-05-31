func xorOperation(n int, start int) (ans int) {
    for i := 0; i < n; i++ {
        ans ^= start + (i << 1)
    }

    return
}