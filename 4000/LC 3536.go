func maxProduct(n int) int {
    a, b := 0, 0
    for n > 0 {
        x := n % 10
        if x >= a {
            a, b = x, a
        } else if x > b {
            b = x
        }
        n /= 10 
    }

    return a * b
}