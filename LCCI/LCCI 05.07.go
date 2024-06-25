func exchangeBits(num int) int {
    for i := 0; i < 31; i += 2 {
        a, b := 1 << i, 1 << (i + 1)
        if num & a == 0 && num & b == 0 || num & a != 0 && num & b != 0 {
            continue
        }
        num ^= a
        num ^= b
    }

    return num
}