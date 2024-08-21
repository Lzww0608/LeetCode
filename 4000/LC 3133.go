func minEnd(n int, x int) int64 {
    ans := x
    t := 1
    n--
    for n > 0 {
        for t & x != 0 {
            t <<= 1
        }
        
        if n & 1 == 1 {
            ans |= t
        }

        t <<= 1
        n >>= 1
    }

    return int64(ans)
}