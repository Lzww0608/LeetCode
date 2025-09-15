func fraction(cont []int) []int {
    n := len(cont)
    a, b := 1, 0
    for i := n - 1; i >= 0; i-- {
        x := cont[i]
        a, b = b, a 
        a += b * x

        g := gcd(a, b)
        a /= g 
        b /= g
    }

    return []int{a, b}
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    
    return a
}