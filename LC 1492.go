func kthFactor(n int, k int) int {
    var i int
    for i = 1; i * i <= n; i++ {
        if n % i == 0 {
            k--
            if k == 0 {
                return i
            }
        }
    }  

    i--
    if i * i == n {
        i--
    }

    for ;i > 0; i-- {
        if n % i == 0 {
            k--
            if k == 0 {
                return n / i
            }
        }
    }



    return -1
}