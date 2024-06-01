func gcdOfStrings(str1 string, str2 string) string {
    if str1 + str2 != str2 + str1 {
        return ""
    }
    
    m, n := len(str1), len(str2)
    d := gcd(m, n)
    return str1[:d]
}

func gcd(x, y int) int {
    for y > 0 {
        x, y = y, x % y
    }
    return x
}
