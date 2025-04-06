func smallestValue(n int) int {
    for n > 5 {
        cur, t := 0, n
        for x := 2; x * x <= t; x++ {
            for t % x == 0 {
                t /= x                                
                cur += x
            }
        }
        if t > 1 {
            cur += t 
        }
        if cur == n {
            break
        }
        n = cur
    }

    return n
}