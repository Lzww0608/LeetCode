func canAliceWin(n int) bool {
    for x := 10; x > 0; x-- {
        if n < x {
            return x & 1 == 1
        }
        n -= x
    }

    return true
}