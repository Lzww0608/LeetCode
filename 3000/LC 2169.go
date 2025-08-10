func countOperations(a int, b int) (ans int) {
    for a != 0 && b != 0 {
        if a > b {
            ans += a / b
            a %= b
        } else {
            ans += b / a 
            b %= a
        }
    }

    return
}