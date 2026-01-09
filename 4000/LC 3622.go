func checkDivisibility(n int) bool {
    sum, prod := 0, 1
    for x := n; x > 0; x /= 10 {
        t := x % 10
        sum += t 
        prod *= t
    }

    return n % (sum + prod) == 0
}