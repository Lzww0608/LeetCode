func isPerfectSquare(num int) bool {
    //newton iteration
    x0 := float64(num)

    for {
        x1 := (x0 + float64(num) / x0) / 2
        if x0 - x1 < 1e-6 {
            return int(x0) * int(x0) == num
        }
        x0 = x1
    }

    return false
}
