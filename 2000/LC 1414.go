func findMinFibonacciNumbers(k int) (ans int) {
    a, b := 1, 1
    for b < k {
        a, b = b, a + b
    }

    for k > 0 {
        ans += k / b
        k -= k / b * b 
        
        for b > k {
            a, b = b - a, a
        }
    }

    return ans
}