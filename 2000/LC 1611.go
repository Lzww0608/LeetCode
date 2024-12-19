func minimumOneBitOperations(n int) (ans int) {
    bits := []uint8{}
    for m := n; m > 0; m >>= 1 {
        bits = append(bits, uint8(m & 1))
    }

    cur := 1
    for i := len(bits) - 1; i >= 0; i-- {
        if x := bits[i]; x == 1 {
            ans += cur * ((1 << (i + 1)) - 1)
            cur = -cur
        }
    }

    return
}