func commonFactors(a int, b int) (ans int) {
    c := min(a, b)
    for i := 1; i <= c; i++ {
        if a % i == 0 && b % i == 0 {
            ans++
        }
    }

    return
}