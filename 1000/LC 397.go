func integerReplacement(n int) (ans int) {
    for n > 1 {
        if n & 1 == 0 {
            n >>= 1
        } else {
            if n & 0x3 == 0x3 && n != 3 {
                n++
            } else {
                n--
            }
        }
        ans++
    }

    return
}
// 011 100 10 1 0
// 011 010 01 00 0
