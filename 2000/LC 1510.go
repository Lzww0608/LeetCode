func winnerSquareGame(n int) bool {

    f := make([]bool, n + 1)

    for i := range f {
        for j := int(math.Sqrt(float64(i))); j > 0; j-- {
            if !f[i - j * j] {
                f[i] = true
                break
            }
        }
    }

    return f[n]
    
}
