func lastInteger(n int64) int64 {
    ans := int64(1)
    for d, cnt := int64(2), 1; n > 1; d <<= 1 {
        t := (n - 1) / 2 
        if cnt & 1 == 1 {
            ans = ans + t * d 
        } else {
            ans = ans - t * d
        }
        cnt ^= 1 
        n -= n / 2
    }

    return ans
}