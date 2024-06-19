func primePalindrome(n int) int {
    for {
        if isPalindrome(n) && isPrime(n) {
            return n
        }
        n++
        s := strconv.Itoa(n)
        t := len(s)
        if n > 11 && t & 1 == 0 {
            n = quickPow(10, t)
        }
    }

    return n
}

func quickPow(a, r int) (res int) {
    res = 1
    for ; r > 0; r >>= 1 {
        if r & 1 == 1 {
            res *= a
        }
        a *= a
    }
    return 
}

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i * i <= n; i++ {
        if n % i == 0 {
            return false
        }
    }

    return true
}

func reverse(n int) (ans int) {
    for n > 0 {
        ans = ans * 10 + n % 10
        n /= 10
    }
    return
}

func isPalindrome(n int) bool {
    return n == reverse(n)
}