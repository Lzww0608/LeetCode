func circularGameLosers(n int, k int) (ans []int) {
    f := make([]bool, n)
    i, j := 1, 0; 
    for !f[j] {
        f[j] = true
        j = (j + i * k) % n 
        i++
    }

    for i := range f {
        if !f[i] {
            ans = append(ans, i + 1)
        }
    }

    return 
}