func sumOfTheDigitsOfHarshadNumber(x int) int {
    t, sum := x, 0
    for t > 0 {
        sum += t % 10
        t /= 10
    }
    if x % sum == 0 {
        return sum
    }

    return -1
}
