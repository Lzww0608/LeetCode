func threeFibonacciSum(n int) bool {
    if n == 2 {
        return true
    }
    a, b, c := 0, 1, 1
    
    for a + b + c < n {
        a, b, c = b, c, a + b 
        if a + b + c == n {
            return true
        }
    }

    return false
}