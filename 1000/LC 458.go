func poorPigs(buckets int, minutesToDie int, minutesToTest int) (ans int) {
    t := minutesToTest / minutesToDie + 1
    for fastPow(t, ans) < buckets {
        ans++
    }

    return 
}

func fastPow(a, r int) int {
    res := 1
    for ; r > 0; r >>= 1 {
        if r & 1 == 1 {
            res *= a
        }
        a *= a
    }

    return res
}