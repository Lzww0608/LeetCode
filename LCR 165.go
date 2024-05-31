func crackNumber(ciphertext int) int {
    s := strconv.Itoa(ciphertext)
    n := len(s)
    a, b := 1, 1
    for i := 1; i < n; i++ {
        tmp := b
        if s[i-1:i+1] >= "10" && s[i-1:i+1] < "26" {
            tmp += a 
        }
        a, b = b, tmp
    }

    return b
}