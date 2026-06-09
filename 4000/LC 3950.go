func consecutiveSetBits(n int) bool {
    last := -1
    cnt := 0
    for n > 0 {
        x := n & 1 
        if x == last && x == 1 {
            cnt++
        }
        last = n & 1
        n >>= 1
    }

    return cnt == 1
}