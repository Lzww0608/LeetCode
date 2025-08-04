func chalkReplacer(chalk []int, k int) int {
    sum := 0
    for _, x := range chalk {
        sum += x
    }

    k %= sum
    cur := 0
    for i, x := range chalk {
        cur += x 
        if cur > k {
            return i
        }
    }

    return -1
}