func subtractProductAndSum(n int) int {
    prod, sum := 1, 0
    for n > 0 {
        t := n % 10
        n /= 10
        prod *= t
        sum += t
    }

    return prod - sum
}