var MOD int = int(1e9 + 7)
func nthMagicalNumber(n int, a int, b int) int {
    a, b = min(a, b), max(a, b)
    c := lcm(a, b) 
    cnt := c / a + c / b - 1
    x, y := n / cnt, n % cnt
    l, r := 0, c
    for l < r {
        mid := l + ((r - l) >> 1)
        if mid / a + mid / b >= y {
            r = mid
        } else {
            l = mid + 1
        }
    }
    return (l + x * c % MOD) % MOD
}

func lcm(a, b int) int {
    ans := a * b
    for b != 0 {
        a, b = b, a % b
    }
    return ans / a
}