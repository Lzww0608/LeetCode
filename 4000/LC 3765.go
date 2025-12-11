func completePrime(num int) bool {
    s := strconv.Itoa(num)
    n := len(s)

    x, m := 0, 1 
    for i := range n {
        x = x * 10 + int(s[i] - '0')
        m *= 10
        if !check(x) {
            return false
        }
    }

    for m /= 10; m > 0; m /= 10 {
        if !check(x % m) {
            return false
        }
    }

    return true
}

func check(x int) bool {
    if x == 1 {
        return false
    }
    for i := 2; i * i <= x; i++ {
        if x % i == 0 {
            return false
        }
    }

    return true
}