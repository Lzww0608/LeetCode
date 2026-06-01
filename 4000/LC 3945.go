func digitFrequencyScore(n int) (ans int) {
    for n > 0 {
        x := n % 10
        ans += x 
        n /= 10
    }

    return
}