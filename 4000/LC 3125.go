func maxNumber(n int64) int64 {
    i := 0
    for (1 << i) <= n {
        i++
    }
    return 1 << (i - 1)  - 1
}