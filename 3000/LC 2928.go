func distributeCandies(n int, limit int) (ans int) {
    for i := 0; i <= limit; i++ {
        for j := 0; j <= limit; j++ {
            if n - i - j <= limit && n - i - j >= 0 {
                ans++
            }
        }
    } 
    return
}
