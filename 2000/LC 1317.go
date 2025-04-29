func getNoZeroIntegers(n int) []int {
    check := func(x int) bool {
        for x > 0 {
            t := x % 10
            if t == 0 {
                return false
            }
            x /= 10
        }
        return true
    }
    for {
        i := rand.Intn(n / 2) + 1
        if check(i) && check(n - i) {
            return []int{i, n - i}
        }
    }

    return nil
}