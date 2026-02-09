func removeZeros(n int64) (res int64) {
    var ans int64
    for n > 0 {
        if n % 10 != 0 {
            ans = ans * 10 + n % 10
        }
        n /= 10
    }

    for ans > 0 {
        res = res * 10 + ans % 10 
        ans /= 10
    }
    
    return
}
