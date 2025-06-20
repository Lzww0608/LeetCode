func countDigits(num int) (ans int) {
    x := num 
    for x > 0 {
        t := x % 10 
        if num % t == 0 {
            ans++
        }
        x /= 10
    }

    return
}