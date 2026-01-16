func gcdOfOddEvenSums(n int) int {
    a, b := 0, 0
    for i := range n {
        a += i * 2 + 1
        b += i * 2 + 2
    }

    return gcd(a, b)
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b 
    }

    return a
}