func lateFee(daysLate []int) (ans int) {
    for _, x := range daysLate {
        if x == 1 {
            ans += x
        } else if x <= 5 {
            ans += x * 2
        } else {
            ans += x * 3
        }
    }

    return
}