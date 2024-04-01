func selfDividingNumbers(left int, right int) (ans []int) {
    check := func(x int) bool {
        n := x
        for x > 0 {
            t := x % 10
            if t == 0 || n % t != 0 {
                return false
            }  
            x /= 10
        }
        return true
    }

    for i := left; i <= right; i++ {
        if check(i) {
            ans = append(ans, i)
        }
    }

    return 
}